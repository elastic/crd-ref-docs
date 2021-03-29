{{- define "type" -}}
{{- $type := . -}}
{{- if rstShouldRenderType $type -}}

.. _{{ $type.Name  }} {{ if $type.IsAlias }}({{ rstRenderTypeLink $type.UnderlyingType  }}) {{ end }}:
{{rstBlankLine}}
{{ $type.Name  }} {{ if $type.IsAlias }}({{ rstRenderTypeLink $type.UnderlyingType  }}) {{ end }}
{{ rstRenderUnderline $type.Name "^" }}

{{ $type.Doc }}

{{ if $type.References -}}
Appears In:
{{- range $type.SortedReferences }}
{{ rstRenderTypeLink . }} 
{{- end }}

{{- end }}


{{ if $type.Members -}}
.. csv-table:: 
   :header: "Field", "Description"
   :widths: 10, 40

{{ if $type.GVK -}}
   {{ rstSpaces 3 }}"``apiVersion`` (string)", "`{{ $type.GVK.Group }}/{{ $type.GVK.Version }}`"
   {{ rstSpaces 3 }}"``kind`` (string)", "`{{ $type.GVK.Kind }}`"
{{ end -}}

{{ range $type.Members -}}
   {{ rstSpaces 3 }}"``{{ .Name  }}`` ({{ rstRenderType .Type }})", {{ template "type_members" . }}
{{ end -}}
{{ end -}}



{{- end -}}
{{- end -}}
