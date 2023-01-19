// Copyright 2022 Dimitrij Drus <dadrus@gmx.de>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package matcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeaderMatcher(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		uc       string
		headers  map[string][]string
		match    map[string]string
		matching bool
	}{
		{
			uc: "match single header",
			headers: map[string][]string{
				"foobar": {"foo", "bar"},
			},
			match:    map[string]string{"foobar": "bar,baz"},
			matching: true,
		},
		{
			uc: "match multiple header",
			headers: map[string][]string{
				"foobar":      {"foo", "bar"},
				"some-header": {"value1", "value2"},
			},
			match: map[string]string{
				"foobar":      "bar,foo",
				"some-header": "value1,val3",
			},
			matching: true,
		},
		{
			uc: "don't match header",
			headers: map[string][]string{
				"foobar":      {"foo", "bar"},
				"some-header": {"value1", "value2"},
			},
			match:    map[string]string{"barfoo": "bar"},
			matching: false,
		},
		{
			uc: "don't match header value",
			headers: map[string][]string{
				"foobar":      {"foo", "bar"},
				"some-header": {"value1", "value2"},
			},
			match:    map[string]string{"foobar": "value1"},
			matching: false,
		},
	} {
		t.Run("case="+tc.uc, func(t *testing.T) {
			matcher := HeaderMatcher(tc.headers)

			// WHEN
			matched := matcher.Match(tc.match)

			// THEN
			assert.Equal(t, tc.matching, matched)
		})
	}
}