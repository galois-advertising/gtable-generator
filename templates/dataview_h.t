#pragma once
// solopointer1202@gmail.com

#include "dataview.h"

namespace {{ .Namespace }} {
class {{ $.Name -}}_traits {
public:
    using insert_raw_t = typename {{.Datasource_name -}}_{{- .Datasource_channel -}}_etraits::update_t;
    using insert_derivative_t = typename {{.Datasource_name -}}_{{- .Datasource_channel -}}_etraits::derivative_t;
    using update_raw_t = typename {{.Datasource_name -}}_{{- .Datasource_channel -}}_etraits::update_t;
    using update_derivative_t = typename {{.Datasource_name -}}_{{- .Datasource_channel -}}_etraits::derivative_t;
    using delete_key_t = typename {{.Datasource_name -}}_{{- .Datasource_channel -}}_etraits::delete_t;
};
typedef galois::gtable::dataview<{{- $.Name -}}_traits> {{ $.Name -}}_dataview;

}
