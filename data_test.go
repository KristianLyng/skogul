/*
 * skogul, data structure tests
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

package skogul_test

import (
	"fmt"
	"github.com/KristianLyng/skogul"
	"github.com/KristianLyng/skogul/parser"
	"testing"
	"time"
)

// Test the String() capability of containers. Unfortunately, this depends
// on marshaling happening in a predictable order, which there is no
// guarantee that it will. But it seems to work ok for now.
func TestString(t *testing.T) {
	want := `{
  "template": {},
  "metrics": [
    {
      "timestamp": "2019-03-15T11:08:02+01:00",
      "metadata": {
        "key": "value"
      },
      "data": {
        "integer": 5
      }
    }
  ]
}`
	b := []byte(want)
	j := parser.JSON{}
	c, err := j.Parse(b)
	if err != nil {
		t.Errorf("JSON.Parse(b) failed: %v", err)
		return
	}
	got := fmt.Sprintf("%s", c)

	if got != want {
		t.Errorf("String() on container failed, wanted %v, got %v", want, got)
	}
}

func TestValidate(t *testing.T) {
	empty := skogul.Container{}
	err := empty.Validate()
	if err == nil {
		t.Errorf("Validate() succeeded on an empty Container")
	}
	got := fmt.Sprintf("%s", err)
	want := "<nil>: Missing metrics[] data"
	if got != want {
		t.Errorf("Validate() expected reason %s, got %s", want, got)
	}

	noMetrics := skogul.Container{}
	noMetrics.Metrics = []skogul.Metric{}
	err = noMetrics.Validate()
	if err == nil {
		t.Errorf("Validate() succeeded on an Container with empty metrics[]")
	}
	got = fmt.Sprintf("%s", err)
	want = "<nil>: Empty metrics[] data"
	if got != want {
		t.Errorf("Validate() expected reason %s, got %s", want, got)
	}

	badMetrics := skogul.Container{}
	metric := skogul.Metric{}
	metric.Data = make(map[string]interface{})
	badMetrics.Metrics = []skogul.Metric{metric}
	err = badMetrics.Validate()
	if err == nil {
		t.Errorf("Validate() succeeded on an Container with empty metrics[]")
	}
	got = fmt.Sprintf("%s", err)
	want = "<nil>: Missing timestamp in both metric and container"
	if got != want {
		t.Errorf("Validate() expected reason %s, got %s", want, got)
	}

	now := time.Now()
	metric.Time = &now
	notimeMetrics := skogul.Container{}
	notimeMetrics.Metrics = []skogul.Metric{metric}
	err = notimeMetrics.Validate()
	if err == nil {
		t.Errorf("Validate() succeeded on an Container with no data")
	}
	got = fmt.Sprintf("%s", err)
	want = "<nil>: Missing data for metric"
	if got != want {
		t.Errorf("Validate() expected reason %s, got %s", want, got)
	}

	metric.Data["test"] = "foo"
	okc := skogul.Container{}
	okc.Metrics = []skogul.Metric{metric}
	err = okc.Validate()
	if err != nil {
		t.Errorf("Validate() failed when it should work: %v", err)
	}
}