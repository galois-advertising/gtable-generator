/* 
This code is generated automatically by gtable-generator, do not modify it. 
solopointer1202@gmail.com
*/
#pragma once
#include <memory>
#include <map>
#include "defines.h"
#include "gtable.h"

// Datasources
{{ range .Datasources -}}
#include "{{.Name -}}.h"
{{ end -}}

// Dataview
{{ range .Dataviews -}}
#include "{{.Name -}}.h"
{{ end -}}
// Datatable 
{{ range .Datatables -}}
#include "{{.Name -}}.h"
{{ end -}}
// Indextable
{{ range .Indextables -}}
#include "{{.Name -}}.h"
{{ end -}}
// Dataupdator
{{ range .Dataupdators -}}
#include "{{.Name -}}.h"
{{ end -}}
// Indexupdator
{{ range .Indexupdators -}}
#include "{{.Name -}}.h"
{{ end }}

namespace {{ .Namespace }} {

class {{ .Handler }}: public galois::gtable::gtable_project {
public:

    {{ .Handler -}}(){};
    virtual ~{{ .Handler -}}(){};
    
    BEGIN_DATASOUECE
    {{ range .Datasources -}}
    DATASOURCE({{- .Name -}}_source)
    {{ end -}}
    END_DATASOURCE
    // Dataview
    {{ range .Dataviews -}}
    DEFINE_HANDLER({{- .Name -}}_dataview, {{ .Name -}}_var);
    {{ end -}}
    // Datatable 
    {{ range .Datatables -}}
    DEFINE_HANDLER({{- .Name -}}_datatable, {{ .Name -}}_var);
    {{ end -}}
    // Indextable
    {{ range .Indextables -}}
    DEFINE_HANDLER({{- .Name -}}_indextable, {{ .Name -}}_var);
    {{ end -}}
private:
    // Dataupdator
    {{ range .Dataupdators -}}
    DEFINE_HANDLER({{- .Name -}}_dataupdator, {{ .Name -}}_var);
    {{ end -}}
    // Indexupdator
    {{ range .Indexupdators -}}
    DEFINE_HANDLER({{- .Name -}}_indexupdator, {{ .Name -}}_var);
    {{ end -}}

    bool setup_dataupdator() {
        {{ range $idv, $dv := .Dataviews -}}
        {{ range $idu, $du := $dv.Dataupdators -}}
        {{ $du.Name -}}_var().set_datatable(&{{ .To -}}_var()); 
        {{ $dv.Name -}}_var().append_dataupdator(
            dynamic_cast<std::remove_reference<decltype({{- $dv.Name -}}_var())>::type::dataupdator_t>(&{{ $du.Name -}}_var()));
        {{ end -}}
        {{ end -}}
        return true;
    }

    bool setup_indexupdator() {
        {{ range $idt, $dt := .Datatables -}}
        {{ range $iiu, $iu := $dt.Indexupdators -}}
        {{ $iu.Name -}}_var().set_indextable(&{{ .To -}}_var()); 
        {{ $dt.Name -}}_var().append_indexupdator(
            dynamic_cast<std::remove_reference<decltype({{- $dt.Name -}}_var())>::type::indexupdator_t>(&{{ $iu.Name -}}_var()));
        {{ end -}}
        {{ end -}}
        return true;
    }
};


}  