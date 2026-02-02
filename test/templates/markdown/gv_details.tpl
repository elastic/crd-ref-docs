{{- define "gvDetails" -}}
{{- $gv := . -}}

## <a id="{{ markdownGroupVersionID $gv | markdownSafeID }}">{{ $gv.GroupVersionString }}</a>

{{ $gv.Doc }}

{{- if index $gv.Markers "special" }}
*Important: This package is special and should be treated differently.*
{{- end }}

{{- if $gv.Kinds  }}
### Resource Types
{{- range $gv.SortedKinds }}
- {{ $gv.TypeForKind . | markdownRenderTypeLink }}
{{- end }}
{{ end }}

{{ range $gv.SortedTypes }}
{{ template "type" . }}
{{ end }}

{{- end -}}
