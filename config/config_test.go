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
package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad_LinkMappingsValidation(t *testing.T) {
	tests := []struct {
		name    string
		yaml    string
		wantErr bool
	}{
		{
			name: "valid mapping",
			yaml: `render:
  linkMappings:
    - url: https://example.com/old
      link: docs-content://new/page.md
      text: New page
`,
		},
		{
			name: "empty url is rejected",
			yaml: `render:
  linkMappings:
    - url: ""
      link: docs-content://new/page.md
      text: New page
`,
			wantErr: true,
		},
		{
			name: "missing link is rejected",
			yaml: `render:
  linkMappings:
    - url: https://example.com/old
      text: New page
`,
			wantErr: true,
		},
		{
			name: "missing text is rejected",
			yaml: `render:
  linkMappings:
    - url: https://example.com/old
      link: docs-content://new/page.md
`,
			wantErr: true,
		},
		{
			name: "no mappings is valid",
			yaml: `render: {}
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "config.yaml")
			require.NoError(t, os.WriteFile(path, []byte(tt.yaml), 0o600))

			_, err := Load(Flags{Config: path})
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
