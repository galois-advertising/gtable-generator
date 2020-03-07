#pragma once

namespace {{.Namespace}} {

template <typename traits>
class {{ .Get_udf }} {
public:
{{- range $icol, $col := .Columns}}
{{- if eq $col.Colume_from "derivative"}}
    template<typename F, typename T>
    static bool parse_{{- $col.Column_name}}(
            const F& {{ $col.Get_from -}},
            T& {{ $col.Column_name -}}) {
        return false;
    }

{{- end}}
{{- end}}
};

}  
