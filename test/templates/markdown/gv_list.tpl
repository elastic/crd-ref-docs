{{- define "gvList" -}}
{{- $groupVersions := . -}}

# API Reference

Here is a template value: `{{ markdownTemplateValue "k1" }}`.

## Packages
{{- range $groupVersions }}
- {{ markdownRenderGVLink . }}
{{- end }}

{{ range $groupVersions }}
{{ template "gvDetails" . }}
{{ end }}

{{- end -}}
