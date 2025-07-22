// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
package renderer

import (
	"testing"

	"github.com/elastic/crd-ref-docs/config"
	"github.com/stretchr/testify/assert"
)

func newTestConfig(t *testing.T) *config.Config {
	t.Helper()
	conf := &config.Config{}
	if err := conf.TemplateKeyValues.Set("foo=bar"); err != nil {
		t.Fatal(err)
	}
	if err := conf.TemplateKeyValues.Set("hello=world"); err != nil {
		t.Fatal(err)
	}
	return conf
}

func TestMarkdownRenderer_TemplateValue(t *testing.T) {
	tests := []struct {
		name     string
		renderer *MarkdownRenderer
		key      string
		want     string
	}{
		{
			name:     "existing key",
			renderer: &MarkdownRenderer{conf: newTestConfig(t)},
			key:      "foo",
			want:     "bar",
		},
		{
			name:     "another existing key",
			renderer: &MarkdownRenderer{conf: newTestConfig(t)},
			key:      "hello",
			want:     "world",
		},
		{
			name:     "missing key",
			renderer: &MarkdownRenderer{conf: newTestConfig(t)},
			key:      "missing",
			want:     "",
		},
		{
			name:     "nil config",
			renderer: &MarkdownRenderer{conf: nil},
			key:      "foo",
			want:     "",
		},
		{
			name:     "nil renderer",
			renderer: nil,
			key:      "foo",
			want:     "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.renderer.TemplateValue(tt.key))
		})
	}
}
