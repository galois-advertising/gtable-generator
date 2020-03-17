/* 
This code is generated automatically by gtable-generator, do not modify it. 
solopointer1202@gmail.com
*/
#pragma once
#include "default_dataupdator.h"
#include "{{- .From -}}.h"
#include "{{- .To -}}.h"

namespace {{ .Namespace }} {
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
    using data_table = {{ .To -}}_datatable;
    using row_t = typename data_table::row_t;

    static bool create_row_tuple(const insert_raw_t&, const insert_derivative_t&, data_table::row_t&);
    template <typename t>
    static data_table::primary_key_t primary_key(const t& row_or_key) {
        return data_table::primary_key(row_or_key);
    }
    static bool update(const row_t& new_tuple, row_t& old_tuple) {

    }
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

using {{ .Name -}}_dataupdator = galois::gtable::default_dataupdator<{{- .From -}}_traits, {{ .Name -}}_traits>;

/*
{{ .From -}}_traits: The type traits of dataview, it's a smaller traits
{{ .Name -}}_traits: The type traits of dataupdator, {{ .Name -}}_traits contains {{ .From -}}_traits. It is a bigger traits.
Why not use {{ .Name -}}_traits directly?
Because of default_dataupdator inherits from idataupdator and idataupdator needs {{ .From -}}_traits.
*/
} 