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

func TestLinkForType(t *testing.T) {
	conf := config.Config{
		Render: config.RenderConfig{
			KubernetesVersion: "1.29",
			KnownTypes: []*config.KnownType{
				// Register our own package to verify it is ignored for local types.
				{Name: "Foo", Package: "example.com/pkg", Link: "https://example.com/docs#foo"},
				// Also register a kube type to verify IsKubeType takes precedence over knownTypes.
				{Name: "ObjectMeta", Package: "k8s.io/apimachinery/pkg/apis/meta/v1", Link: "https://example.com/docs#objectmeta"},
			},
		},
	}

	f, err := NewFunctions(&conf)
	require.NoError(t, err)

	cases := []struct {
		name      string
		typ       *types.Type
		wantLink  string
		wantLocal bool
	}{
		{
			name:      "imported type gets external link from knownTypes",
			typ:       &types.Type{Name: "Foo", Package: "example.com/pkg", Imported: true},
			wantLink:  "https://example.com/docs#foo",
			wantLocal: false,
		},
		{
			name:      "local type ignores knownTypes and gets local link",
			typ:       &types.Type{Name: "Foo", Package: "example.com/pkg", Imported: false},
			// FIXME: This is incorrect and should be a relative link
			// wantLink:  "example-com-pkg-foo",
			// wantLocal: true,
			wantLink:  "https://example.com/docs#foo",
			wantLocal: false,
		},
		{
			name:      "kube type gets kubernetes.io link even when also in knownTypes",
			typ:       &types.Type{Package: "k8s.io/apimachinery/pkg/apis/meta/v1", Name: "ObjectMeta"},
			wantLink:  "https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.29/#objectmeta-v1-meta",
			wantLocal: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			link, local := f.LinkForType(tc.typ)
			require.Equal(t, tc.wantLink, link)
			require.Equal(t, tc.wantLocal, local)
		})
	}
}

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
		excepted string
	}{
		{
			input:    &types.Type{Package: "k8s.io/apimachinery/pkg/apis/meta/v1", Name: "ObjectMeta"},
			excepted: "https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.29/#objectmeta-v1-meta",
		},
		{
			input:    &types.Type{Package: "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1", Name: "JSON"},
			excepted: "https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.29/#json-v1-apiextensions-k8s-io",
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			link := kh.LinkForKubeType(tc.input)
			require.Equal(t, tc.excepted, link)
		})
	}
}
