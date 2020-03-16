/* 
This code is generated automatically by gtable-generator, do not modify it. 
solopointer1202@gmail.com
*/
#include "{{.Name -}}.h"

namespace {{.Namespace}} {

std::ostream & operator<< (std::ostream& os, const {{.Name}}_row_tuple& tuple) {
{{- range $index, $element := .Columns -}}
{{- if $element.IsArray }}
    if (tuple.is_set_{{- $element.Column_name }}()) {
        os << "|{{ $element.Column_name -}}: [";
        for (auto& i : tuple.{{- $element.Column_name -}}()) {
            os << i << ",";
        }
        os << "]";
    } else {
        os << "|{{ $element.Column_name -}}: [-]";
    }
{{- else }}
    if (tuple.is_set_{{- $element.Column_name -}}()) {
        os << "|{{ $element.Column_name -}}: " << tuple.{{- $element.Column_name -}}();
    } else {
        os << "|{{ $element.Column_name -}}: -";
    }
{{- end }}
{{- end }}
    return os;
}

}
