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
/*
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
*/
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
/*
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
*/
    bool dataupdator_linkto_dataview() {
/*
        {{ range $idv, $dv := .Dataviews -}}
        {{ range $idp, $du := $dv.Dataupdators -}}
            {{- $dv.Name -}}_dataview_var()->append_dataupdator({{ $du.Name -}}_var());
        {{ end -}}
        {{ end -}}
*/
        return true;
    }
};


}  