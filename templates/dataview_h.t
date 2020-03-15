#pragma once
// solopointer1202@gmail.com

#include "dataview.h"

namespace {{ .Namespace }} {
class {{ $.Name -}}_traits {
public:
    using insert_raw_t = typename {{.DatasourceName -}}_{{- .DatasourceChannel -}}_etraits::update_t;
    using insert_derivative_t = typename {{.DatasourceName -}}_{{- .DatasourceChannel -}}_etraits::derivative_t;
    using update_raw_t = typename {{.DatasourceName -}}_{{- .DatasourceChannel -}}_etraits::update_t;
    using update_derivative_t = typename {{.DatasourceName -}}_{{- .DatasourceChannel -}}_etraits::derivative_t;
    using delete_key_t = typename {{.DatasourceName -}}_{{- .DatasourceChannel -}}_etraits::delete_t;
    static const char* name(){return "{{ .Name }}";}
};
typedef galois::gtable::dataview<{{- $.Name -}}_traits> {{ $.Name -}}_dataview;

}
