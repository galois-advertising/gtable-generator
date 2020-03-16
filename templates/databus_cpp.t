/* 
This code is generated automatically by gtable-generator, do not modify it. 
solopointer1202@gmail.com
*/

#include "log.h"
#include <type_traits>
#include "project.h"

#include "{{.Name -}}.h"
{{- range $idv, $dv := .Dataviews }}
{{- if $dv.HasUDF }}
#include "{{- $dv.GetUDF }}.h"
{{- end }}
{{- end }}

namespace {{.Namespace}} {

class {{.Handler}};

{{- range $idv, $dv := .Dataviews}}


int {{$dv.DatasourceName -}}_{{- $dv.DatasourceChannel -}}_callbacks::insert({{$dv.Handler -}}* env, const galois::gformat::pack_header_t& header,
    const typename {{$dv.DatasourceName -}}_{{- $dv.DatasourceChannel -}}_etraits::update_t& raw) {
    typename {{$dv.DatasourceName -}}_{{- $dv.DatasourceChannel -}}_etraits::derivative_t derivative;

{{- range $icol, $col := .Columns}}

{{- if eq $col.IsDerivative "derivative"}}
    std::remove_reference<decltype(derivative.{{- $col.Column_name -}}())>::type temp_{{- $col.Column_name -}};
    if (!{{- $dv.GetUDF -}}<int>::parse_{{- $col.Column_name -}}(raw.{{- $col.Parse_from -}}(), temp_{{- $col.Column_name -}})) {
        FATAL("Parse %s failed!", "{{- $col.Column_name -}}");
        return -1;
    } else {
       derivative.set_{{- $col.Column_name -}}(temp_{{- $col.Column_name -}});
    }
{{- end}}

{{- end}}
    if (!env->{{- $dv.Name -}}_var().notify_insert(raw, derivative)) {
        FATAL("{{- $dv.Name -}}_var().notify_insert failed", "");
        return -1;
    }
    return 0;
};

int {{$dv.DatasourceName -}}_{{- $dv.DatasourceChannel -}}_callbacks::del({{$dv.Handler -}}* env, const galois::gformat::pack_header_t& header,
    const typename {{$dv.DatasourceName -}}_{{- $dv.DatasourceChannel -}}_etraits::delete_t& raw) {
    return 0;
};

int {{$dv.DatasourceName -}}_{{- $dv.DatasourceChannel -}}_callbacks::update({{$dv.Handler -}}* env, const galois::gformat::pack_header_t& header,
    const typename {{$dv.DatasourceName -}}_{{- $dv.DatasourceChannel -}}_etraits::update_t& raw) {
    return 0;
};

{{end}}


}; // End of namespace