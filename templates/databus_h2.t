typedef {{.Handler -}}* {{.Datasource_name -}}_gtable_env_t;
class {{.Datasource_name -}}_traits :
    public galois::gdatabus::default_traits<
        {{.Datasource_name -}}_gtable_env_t
    > {
public:
    typedef {{.Handler}} handle_t;

{{- range .Dataview_typedef_list}}
#include "{{. -}}"
{{- end}}

};

typedef galois::gtable::databus_datasource<{{.Datasource_name -}}_traits> {{.Datasource_name}};
}