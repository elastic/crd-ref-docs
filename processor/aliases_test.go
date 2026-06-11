package processor

import (
	"testing"

	"github.com/elastic/crd-ref-docs/config"
	"github.com/stretchr/testify/require"
)

func TestHasCaseIgnore(t *testing.T) {
	testCases := []struct {
		name    string
		options []string
		want    bool
	}{
		{name: "empty", options: nil, want: false},
		{name: "name only", options: []string{"groupWait"}, want: false},
		{name: "case ignore", options: []string{"groupWait", "case:ignore"}, want: true},
		{name: "case ignore with spaces", options: []string{"groupWait", " case:ignore "}, want: true},
		{name: "case ignore among options", options: []string{"groupWait", "omitempty", "case:ignore"}, want: true},
		{name: "case strict", options: []string{"groupWait", "case:strict"}, want: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, hasCaseIgnore(tc.options))
		})
	}
}

func TestToSnakeCase(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{in: "groupWait", want: "group_wait"},
		{in: "apiURL", want: "api_url"},
		{in: "URLApi", want: "url_api"},
		{in: "group_wait", want: "group_wait"},
		{in: "HTTPSProxy", want: "https_proxy"},
		{in: "GroupWait", want: "group_wait"},
		{in: "v1Beta", want: "v1_beta"},
		{in: "enabled", want: "enabled"},
		{in: "x", want: "x"},
	}

	for _, tc := range testCases {
		t.Run(tc.in, func(t *testing.T) {
			require.Equal(t, tc.want, toSnakeCase(tc.in))
		})
	}
}

func TestToCamelCase(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{in: "group_wait", want: "groupWait"},
		{in: "api_url", want: "apiUrl"},
		{in: "groupWait", want: "groupWait"},
		{in: "enabled", want: "enabled"},
		{in: "x", want: "x"},
	}

	for _, tc := range testCases {
		t.Run(tc.in, func(t *testing.T) {
			require.Equal(t, tc.want, toCamelCase(tc.in))
		})
	}
}

func TestDeriveAliases(t *testing.T) {
	testCases := []struct {
		name          string
		conventions   []config.NamingConvention
		canonicalName string
		want          []string
	}{
		{
			name:          "no conventions",
			conventions:   nil,
			canonicalName: "groupWait",
			want:          nil,
		},
		{
			name:          "snake alias for camel field",
			conventions:   []config.NamingConvention{config.NamingConventionSnakeCase},
			canonicalName: "groupWait",
			want:          []string{"group_wait"},
		},
		{
			name:          "camel alias for snake field",
			conventions:   []config.NamingConvention{config.NamingConventionCamelCase},
			canonicalName: "group_wait",
			want:          []string{"groupWait"},
		},
		{
			name:          "identical derivation omitted",
			conventions:   []config.NamingConvention{config.NamingConventionCamelCase, config.NamingConventionSnakeCase},
			canonicalName: "groupWait",
			want:          []string{"group_wait"},
		},
		{
			name:          "single word yields no aliases",
			conventions:   []config.NamingConvention{config.NamingConventionCamelCase, config.NamingConventionSnakeCase},
			canonicalName: "enabled",
			want:          nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := &processor{compiledConfig: &compiledConfig{caseIgnoreAliases: tc.conventions}}
			require.Equal(t, tc.want, p.deriveAliases(tc.canonicalName))
		})
	}
}
