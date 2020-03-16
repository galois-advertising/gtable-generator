/* 
This code is generated automatically by gtable-generator, do not modify it. 
solopointer1202@gmail.com
*/
#pragma once
#include <exception>
#include "databus_datasource.h"
#include "loader.h"

{{- .Cppcode -}}

namespace {{.Namespace}} {
class {{.Handler}};

{{- range .Dataviews}}

typedef {{.Handler -}}* {{.DatasourceName -}}_{{- .DatasourceChannel -}}_gtable_env_t;
typedef galois::gdatabus::event_traits_t<{{- .DatasourceChannel -}}> {{.DatasourceName -}}_{{- .DatasourceChannel -}}_etraits;
class {{.DatasourceName -}}_{{- .DatasourceChannel -}}_callbacks : public
    galois::gdatabus::schema_callbacks<
    {{.DatasourceName -}}_{{- .DatasourceChannel -}}_gtable_env_t,
    {{.DatasourceName -}}_{{- .DatasourceChannel -}}_etraits
> {
public:
    static int insert({{.Handler -}}* env, const galois::gformat::pack_header_t& header,
        const typename {{.DatasourceName -}}_{{- .DatasourceChannel -}}_etraits::update_t& data);

    static int del({{.Handler -}}* env, const galois::gformat::pack_header_t& header,
        const typename {{.DatasourceName -}}_{{- .DatasourceChannel -}}_etraits::delete_t& data);

    static int update({{.Handler -}}* env, const galois::gformat::pack_header_t& header,
        const typename {{.DatasourceName -}}_{{- .DatasourceChannel -}}_etraits::update_t& data);
};

{{- end}}

class {{.Name -}}_schema : 
    public galois::gdatabus::default_traits<{{.Handler -}}*> {
public:
    typedef {{.Handler -}}* gtable_env;
{{- range .Dataviews}}
    typedef {{.DatasourceName -}}_{{- .DatasourceChannel -}}_etraits {{.DatasourceChannel -}};
    typedef {{.DatasourceName -}}_{{- .DatasourceChannel -}}_callbacks {{.DatasourceChannel -}}_callbacks;
{{end}}
};


typedef galois::gtable::databus_datasource<{{.Name -}}_schema> {{.Name -}}_source;

}; // End of namespace
