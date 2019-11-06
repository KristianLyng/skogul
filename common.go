/*
 * skogul, common trivialities
 *
 * Copyright (c) 2019 Telenor Norge AS
 * Author(s):
 *  - Kristian Lyngstøl <kly@kly.no>
 *
 * This library is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.
 *
 * This library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with this library; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA
 * 02110-1301  USA
 */

package skogul

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
	"time"
)

/*
Handler determines what a receiver will do with data received. It requires
a parser to interperet the raw data, 0 or more transformers to mutate
Containers, and a sender to call after data is parsed and mutated and ready
to be dealt with.

Whenever a new Container is created, it should pass that to a Handler, not
directly to a Sender. This goes for artificially created data too, e.g. if
a sender wants to emit statistics. This ensures that transformers can be
used in the future.

To make it configurable, a HandlerRef should be used.
*/
type Handler struct {
	Parser       Parser
	Transformers []Transformer
	Sender       Sender
}

// Parser is the interface for parsing arbitrary data into a Container
type Parser interface {
	Parse(data []byte) (*Container, error)
}

/*
Sender accepts data through Send() - and "sends it off". The canonical
sender is one that implements a storage backend or outgoing API. E.g.:
accept data, send to influx.

Senders are not allowed to modify the Container - there could be multiple
goroutines running with same Container. If modification is required, the
Sender needs to take a copy.

A sender should assume that the container has been validated, and is
non-null. Slightly counter to common sense, it is NOT recommended to
verify the input data again, since multiple senders are likely chained
and will thus likely redo the same verifications.

Senders that pass data off to other senders should use a SenderRef instead,
to facilitate configuration.
*/
type Sender interface {
	Send(c *Container) error
}

/*
Transformer mutates a collection before it is passed to a sender. Transformers
should be very fast, but are the only means to modifying the data.
*/
type Transformer interface {
	Transform(c *Container) error
}

/*
Receiver is how we get data. Receivers are responsible for getting raw data and the
outer boundaries of a Container, but should explicitly avoid parsing raw data.
This ensures that how data is transported is not bound by how it is parsed.
*/
type Receiver interface {
	Start() error
}

/*
Verifier is an *optional* interface for senders and receivers. If
implemented, the configuration engine will issue Verify() after all
configuration is parsed. The sender/receiver should never modify state upon
Verify(), but should simply check that internal state is usable.
*/
type Verifier interface {
	Verify() error
}

/*
Error is a typical skogul error. All Skogul functions should provide Source
and Reason. I'm not entirely sure why, except that it allows chaining errors?

If the Next field is provided, error messages will recurse to the bottom, thus
propagating errors from the bottom and up.
*/
type Error struct {
	Reason string
	Source string
	Next   error
}

/*
SenderRef is a reference to a named sender. This is required to allow
references to be resolved after all senders are loaded. Wherever a
Sender is loaded from configuration, a SenderRef should be used in its
place. The maintenance of the sender is handled in the configuration
system.
*/
type SenderRef struct {
	S    Sender
	Name string
}

// HandlerRef references a named handler. Used whenever a handler is
// defined by configuration.
type HandlerRef struct {
	H    *Handler
	Name string
}

// Error for use in regular error messages. Also outputs to log.Print().
// Will also include e.Next, if present.
func (e Error) Error() string {
	src := "<nil>"
	if e.Source != "" {
		src = e.Source
	}
	tail := ""
	if e.Next != nil {
		tail = fmt.Sprint(": ", e.Next.Error())
	}
	return fmt.Sprintf("%s: %s%s", src, e.Reason, tail)
}

// Container returns a skogul container representing the error
func (e Error) Container() Container {
	c := Container{}
	now := time.Now()
	c.Metrics = make([]*Metric, 1)
	m := Metric{}
	m.Metadata = make(map[string]interface{})
	m.Data = make(map[string]interface{})
	m.Time = &now
	m.Metadata["source"] = e.Source
	m.Data["reason"] = e.Reason
	m.Data["description"] = e.Error()
	c.Metrics[0] = &m
	return c
}

// Transform runs all available transformers in order and enforce any rules
func (h *Handler) Transform(c *Container) error {
	for _, t := range h.Transformers {
		if err := t.Transform(c); err != nil {
			return err
		}
	}
	return nil
}

// Handle parses the byte array using the configured parser, issues
// transformers and sends the data off.
func (h *Handler) Handle(b []byte) error {
	c, err := h.Parser.Parse(b)
	if err != nil {
		return Error{Source: "handler", Reason: "parsing bytestream failed", Next: err}
	}
	if err = c.Validate(); err != nil {
		return Error{Source: "handler", Reason: "validating metrics failed", Next: err}
	}
	if err = h.Transform(c); err != nil {
		return Error{Source: "handler", Reason: "transforming metrics failed", Next: err}
	}
	if err = h.Sender.Send(c); err != nil {
		return Error{Source: "handler", Reason: "sending metrics failed", Next: err}
	}
	return nil
}

// TransformAndSend transforms the already parsed container and sends the
// data off.
func (h *Handler) TransformAndSend(c *Container) error {
	if err := h.Transform(c); err != nil {
		return Error{Source: "handler", Reason: "transforming metrics failed", Next: err}
	}
	if err := h.Sender.Send(c); err != nil {
		return Error{Source: "handler", Reason: "sending metrics failed", Next: err}
	}
	return nil
}

// Verify the basic integrity of a handler. Quite shallow.
func (h Handler) Verify() error {
	if h.Parser == nil {
		return Error{Reason: "Missing parser for Handler"}
	}
	if h.Transformers == nil {
		return Error{Reason: "Missing parser for Handler"}
	}
	for _, t := range h.Transformers {
		if t == nil {
			return Error{Reason: "nil-transformer for Handler"}
		}
	}
	if h.Sender == nil {
		return Error{Reason: "Missing parser for Handler"}
	}
	return nil
}

// Logger returns a logrus.Entry pre-populated with standard Skogul fields.
// sourceType is the typical family of the code/module:
// sender/receiver/parser/transformer/core, while sourceName is the local
// implementation.
func Logger(sourceType, sourceName string) *logrus.Entry {
	return logrus.WithField("source", sourceType).WithField(sourceType, sourceName)
}

// AssertErrors counts the number of assert errors
var AssertErrors int

// Assert panics if x is false, useful mainly for doing error-checking for
// "impossible scenarios" we can't really handle anyway.
//
// Keep in mind that net/http steals panics, but you can check
// AssertErrors, which is incremented with each assert error encountered.
func Assert(x bool, v ...interface{}) {
	if !x {
		out := "assertion failed"
		pc, file, line, ok := runtime.Caller(1)
		if ok {
			out = fmt.Sprintf("%X:%s:%d assertion failed", pc, file, line)
		}
		AssertErrors++
		panic(fmt.Sprintf("%s %s", out, fmt.Sprint(v...)))
	}
}

// ExtractNestedObject extracts an object from a nested object structure. All intermediate objects has to map[string]interface{}
func ExtractNestedObject(object map[string]interface{}, keys []string) (map[string]interface{}, error) {
	if len(keys) == 1 {
		return object, nil
	}

	next, ok := object[keys[0]].(map[string]interface{})

	if !ok {
		return nil, Error{Reason: "Failed to cast nested object to map[string]interface{}"}
	}

	return ExtractNestedObject(next, keys[1:])
}
