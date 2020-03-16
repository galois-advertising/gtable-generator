/* 
This code is generated automatically by gtable-generator, do not modify it. 
solopointer1202@gmail.com
*/
#pragma once
#include "default_indexupdator.h"
#include "{{- .From -}}.h"
#include "{{- .To -}}.h"

namespace {{ .Namespace }} {

template <typename datatable_traits>
class {{ .Name -}}_traits {
public:
    typedef {{ .Handler }} THandle;
    using row_t = typename datatable_traits::row_t;
    using index_table = {{ .To -}}_indextable;
};

using {{ .Name -}}_indexupdator = galois::gtable::default_indexupdator<{{- .From -}}_traits, {{ .Name -}}_traits>;

}