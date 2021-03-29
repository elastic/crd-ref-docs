{{- define "gvList" -}}
{{- $groupVersions := . -}}

API Reference
=============

Packages
--------
{{ rstBlankLine }}
{{- range $groupVersions }}
- {{ rstRenderGVLink . }}
{{- end }}

{{ range $groupVersions }}
{{ template "gvDetails" . }}
{{ end }}

{{- end -}}
