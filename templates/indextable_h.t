/* 
This code is generated automatically by gtable-generator, do not modify it. 
solopointer1202@gmail.com
*/
#pragma once
#include "indextable.h"
#include "{{- .OnTable -}}.h"

namespace {{ .Namespace }} {

template<typename datatable_traits>
class {{ .Name -}}_traits_helper {
public:
    using row_t = typename datatable_traits::row_t;
    using index_key = {{ .KeyType }};
    using primary_key_t = typename datatable_traits::primary_key_t;
    static const index_key& make_index_key(const row_t& row) {
        return row.{{ .OnColumn }}();
    }

};
using {{ .Name -}}_traits = {{ .Name -}}_traits_helper<{{- .OnTable -}}_traits>;
using {{ .Name -}}_indextable = galois::gtable::indextable<{{ .Name -}}_traits>;
}