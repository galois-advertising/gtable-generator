#pragma once
// solopointer1202@gmail.com
#include "default_dataupdator.h"
#include "{{- .From -}}.h"
#include "{{- .To -}}.h"

namespace {{ .Namespace }}{
class {{ .Handler }};

template <typename dataview_traits>
class {{ .Name -}}_traits {
public:
    typedef {{ .Handler }} THandle;
    using insert_raw_t = typename dataview_traits::insert_raw_t;
    using insert_derivative_t = typename dataview_traits::insert_derivative_t;
    using update_raw_t = typename dataview_traits::update_raw_t;
    using update_derivative_t = typename dataview_traits::update_derivative_t;
    using delete_key_t = typename dataview_traits::delete_key_t;
    using data_table = {{ .To }}_datatable;

    static bool create_row_tuple(const insert_raw_t&, const insert_derivative_t&, data_table::row_t&);
    static bool make_primary_key(const insert_raw_t&, const insert_derivative_t&, data_table::primary_key_t& pk);
};

template <typename dataview_traits>
bool {{ .Name -}}_traits<dataview_traits>::create_row_tuple(const insert_raw_t& original, 
    const insert_derivative_t& derivative, 
    data_table::row_t& row) {
    {{- range $idx, $col := .To_datatable.Columns }}
        {{- if or $col.IsBasic $col.IsString }}
            {{- if $col.IsPrimarykey }}
    row.set_{{- $col.Column_name }}({{- $col.IsDerivative -}}.key().{{- $col.Column_name -}}());
            {{- else }}
    row.set_{{- $col.Column_name }}({{- $col.IsDerivative -}}.{{- $col.Column_name -}}());
            {{- end}}
        {{- else if $col.IsArray }}
    {
        int {{ $col.Column_name -}}_max_len = {{ $col.Length }};
        for (auto& i : {{ $col.IsDerivative -}}.{{- $col.Column_name -}}()) {
            if ({{- $col.Column_name -}}_max_len -- <= 0) {
                break;
            }
            row.append_{{- $col.Column_name }}(i);
        }
    }
        {{- end }}
    {{- end }}
    return true;
}

template <typename dataview_traits>
bool {{ .Name -}}_traits<dataview_traits>::make_primary_key(const insert_raw_t&, 
    const insert_derivative_t&, 
    data_table::primary_key_t& pk) {
    return true;
}

using {{ .Name -}}_dataupdator = galois::gtable::default_dataupdator<{{- .From -}}_traits, {{ .Name -}}_traits>;

} 