typedef {{.Handler -}}* {{.DatasourceName -}}_gtable_env_t;
class {{.DatasourceName -}}_traits :
    public galois::gdatabus::default_traits<
        {{.DatasourceName -}}_gtable_env_t
    > {
public:
    typedef {{.Handler}} handle_t;

{{- range .Dataview_typedef_list}}
#include "{{. -}}"
{{- end}}

};

typedef galois::gtable::databus_datasource<{{.DatasourceName -}}_traits> {{.DatasourceName}};
}