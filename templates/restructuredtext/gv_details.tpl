{{- define "gvDetails" -}}
{{- $gv := . -}}

.. _{{ $gv.GroupVersionString }}:
{{ rstBlankLine }}
{{ $gv.GroupVersionString }}
{{ rstRenderUnderline $gv.GroupVersionString "-" }}

{{ $gv.Doc }}

{{- if $gv.Kinds  }}
Resource Types
**************
{{ rstBlankLine }}
{{- range $gv.SortedKinds }}
- {{ $gv.TypeForKind . | rstRenderTypeLink }}
{{- end }}
{{ end }}

{{ range $gv.SortedTypes }}
{{ template "type" . }}
{{ end }}

{{- end -}}
