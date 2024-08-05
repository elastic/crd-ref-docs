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
	"github.com/elastic/crd-ref-docs/types"
	"github.com/stretchr/testify/require"
)

func TestKubernetesHelper(t *testing.T) {
	conf := config.Config{
		Render: config.RenderConfig{
			KubernetesVersion: "1.29",
		},
	}

	kh, err := newKubernetesHelper(&conf)
	require.NoError(t, err)

	cases := []struct {
		input    *types.Type
		expected string
	}{
		{
			input:    &types.Type{Package: "k8s.io/apimachinery/pkg/apis/meta/v1", Name: "ObjectMeta"},
			expected: "https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.29/#objectmeta-v1-meta",
		},
		{
			input:    &types.Type{Package: "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1", Name: "JSON"},
			expected: "https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.29/#json-v1-apiextensions-k8s-io",
		},
		{
			input:    &types.Type{Package: "k8s.io/apimachinery/pkg/api/resource", Name: "Quantity"},
			expected: "https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.29/#quantity-resource-core",
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			link := kh.LinkForKubeType(tc.input)
			require.Equal(t, tc.expected, link)
		})
	}
}
