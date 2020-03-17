/* 
This code is generated automatically by gtable-generator, do not modify it. 
solopointer1202@gmail.com
*/
#pragma once
#include <vector>
#include <ostream>
#include <sstream>
#include "memory_datatable.h"
#include "defines.h"


namespace {{.Namespace}} {

struct {{.Name}}_row_tuple {
{{- range $index, $element := .Columns -}}
{{- if $element.IsBasic }}
    NORMAL_COLUMN({{ $element.Column_kind.Type}}, {{$element.Column_name }})
{{- else if $element.IsString }}
    NORMAL_COLUMN(std::string, {{$element.Column_name }})
{{- else if $element.IsArray }}
    ARRAY_COLUMN({{ $element.Column_kind.Type}}, {{$element.Column_name }})
{{- else }}
    #error "Column type not support [{{$element.Column_name }}]"
{{- end }}
{{- end }}
};

std::ostream & operator<< (std::ostream& , const {{.Name}}_row_tuple& );

class {{.Name}}_traits {
public:
    using row_t = {{.Name}}_row_tuple;
    using primary_key_t = {{.Primary_key.Type}};
    static bool merge_row(const row_t& new_tuple, row_t& old_tuple) {

#ifdef _DEBUG
        std::stringstream ss;
#endif
{{- range $index, $element := .Columns -}}
{{- if $element.IsBasic }}
    {{- if eq $element.Column_kind.Type "double"}}
        if (new_tuple.{{- $element.Column_name }}() < -10e-4 && new_tuple.{{- $element.Column_name }}() > 10e-4 ) {
            old_tuple.set_{{- $element.Column_name -}}(new_tuple.{{- $element.Column_name }}());
#ifdef _DEBUG
            ss << "{{- $element.Column_name -}}:1 ";
#endif
        }
    {{- else}}
        if (new_tuple.{{- $element.Column_name }}() == 0) {
            old_tuple.set_{{- $element.Column_name -}}(new_tuple.{{- $element.Column_name }}());
#ifdef _DEBUG
            ss << "{{- $element.Column_name -}}:1 ";
#endif
        }
    {{- end}}
{{- else if $element.IsString }}
        if (new_tuple.{{- $element.Column_name }}().size() != 0) {
            old_tuple.set_{{- $element.Column_name -}}(new_tuple.{{- $element.Column_name }}());
#ifdef _DEBUG
            ss << "{{- $element.Column_name -}}:1 ";
#endif
        }
{{- else if $element.IsArray }}
        if (new_tuple.{{$element.Column_name }}().size() != 0) {
            old_tuple.set_{{- $element.Column_name -}}(new_tuple.{{$element.Column_name }}());
#ifdef _DEBUG
            ss << "{{- $element.Column_name -}}:1 ";
#endif
        }
{{- else }}
    #error "Column type not support [{{$element.Column_name }}]"
{{- end }}
{{- end }}
#ifdef _DEBUG
        DEBUG("merge_column: %s", ss.str().c_str());
#endif
        return true;
    }
    template<typename t>
    static primary_key_t primary_key(const t& row_or_key) {
        primary_key_t pk = 0;
        
{{- with $pk := .Primary_key }}
{{- range $idx, $key := $pk.Keys }}
        pk ^= row_or_key.{{- $key }}() << {{ $idx }} ;
{{- end}}
{{- end}}
        return pk;
    }
};
using {{.Name}}_datatable = galois::gtable::memory_datatable<{{.Name}}_traits>;

}