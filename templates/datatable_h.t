#pragma once
// solopointer1202@gmail.com
#include <vector>
#include <ostream>
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
    {{.Primary_key.Type}} primary_key() const {
        {{.Primary_key.Type}} pk = 0;
        
{{- with $pk := .Primary_key }}
{{- range $idx, $key := $pk.Keys }}
        pk ^= {{ $key }}() << {{ $idx }} ;
{{- end}}
{{- end}}
        return pk;
    }
};

std::ostream & operator<< (std::ostream& , const {{.Name}}_row_tuple& );

class {{.Name}}_traits {
public:
    using row_t = {{.Name}}_row_tuple;
    using primary_key_t = {{.Primary_key.Type}};
    static bool update(const row_t& tuple, row_t& old) {
        return true;
    }
};
using {{.Name}}_datatable = galois::gtable::memory_datatable<{{.Name}}_traits>;

}