/*
 * skogul, switch transformer
 *
 * Copyright (c) 2019 Telenor Norge AS
 * Author(s):
 *  - Håkon Solbjørg <hakon.solbjorg@telenor.com>
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

package transformer

import (
	"github.com/telenornms/skogul"
)

// Case requires a field ("when") and a value ("is") to match
// for the set of transformers to run
type Case struct {
	When         string   `doc:"Used as a conditional statement on a field"`
	Is           string   `doc:"Used for the specific value (string) of the stated metadata field"`
	Transformers []string `doc:"The transformers to run when the defined conditional is true"`
}

// Switch is a wrapper for a list of cases
type Switch struct {
	Cases []Case `doc:"A list of switch cases "`
}

var logger = skogul.Logger("transformer", "switch")

// Transform checks the cases and applies the matching transformers
func (sw *Switch) Transform(c *skogul.Container) error {
	for _, cas := range sw.Cases {
		field := cas.When
		condition := cas.Is

		for _, metric := range c.Metrics {
			metadataField, ok := metric.Metadata[field].(string)
			if !ok {
				logger.WithField("field", field).Warn("Cast to string for value of metadata field failed")
				continue
			}

			if metadataField == condition {
				for _, wantedTransformerName := range cas.Transformers {
					logger = logger.WithField("wantedTransformer", wantedTransformerName)
					for _, availableTransformer := range skogul.TransformerMap {
						logger = logger.WithField("actualTransformer", availableTransformer.Name)
						// HOW LOW CAN WE GO
						// OH MY GOD
						if wantedTransformerName == availableTransformer.Name {
							logger.Tracef("Transforming with '%s'", availableTransformer.Name)
							(*availableTransformer.T).Transform(c)
							// The name check will never match multiple names, so as soon
							// as we find a match we can break out of this loop
							// and continue processing wanted transformers
							break
						}
					}
				}
			}
		}
	}

	return nil
}