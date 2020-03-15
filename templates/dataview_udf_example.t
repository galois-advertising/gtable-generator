#pragma once

namespace {{.Namespace}} {

template <typename traits>
class {{ .GetUDF }} {
public:
{{- range $icol, $col := .Columns}}
{{- if eq $col.IsDerivative "derivative"}}
    template<typename F, typename T>
    static bool parse_{{- $col.Column_name}}(
            const F& {{ $col.Parse_from -}}_in,
            T& {{ $col.Column_name -}}_out) {
        return false;
    }

{{- end}}
{{- end}}
};

}  
