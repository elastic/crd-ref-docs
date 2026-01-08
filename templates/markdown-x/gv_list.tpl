{{- define "gvList" -}}
{{- $groupVersions := . -}}

<head>
  <meta name="docsearch:indexPrefix" content="reference-doc" />
</head>

# API Reference

## Packages
{{- range $groupVersions }}
- {{ markdownRenderGVLink . }}
{{- end }}

{{ range $groupVersions }}
{{ template "gvDetails" . }}
{{ end }}

{{- end -}}
