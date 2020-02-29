#pragma once
#include "memory_datatable.h"
{{- range .Include_dataview_headers}}
#include "{{. -}}"
{{- end}}

{{.Namespace}} {

class {{.Name}}_basic_tuple;
class {{.Name}}_object_tuple;
class {{.Name}}_object_tuple_ref;

class {{.Name}}_schema {
public:
    typedef {{.Name}}_basic_tuple TBasicTuple;
    typedef {{.Name}}_object_tuple TObjectTuple;
    typedef {{.Primary_key.Type}} TPrimaryKey;
    typedef ${basic_pool}          TBasicPool;
    typedef DataPool<TBasicPool>   TDataPool;
    typedef TBasicPool::TVaddr  TVaddr;

    static const char* PRIMARY_KEY_FIELDS;
    const static uint32_t VAR_POOL_NUM = ${var_pool_num};
    static void make_basic_tuple(const {{.Name}}_object_tuple &tuple, {{.Name}}_basic_tuple* basic);
    static bool is_valid_object_tuple(const {{.Name}}_object_tuple &tuple);
    static void make_object_tuple(const {{.Name}}_object_tuple_ref& ref, {{.Name}}_object_tuple* tuple);
};

struct {{.Name}}_basic_tuple {

{{ range $index, $element := .Columns -}}
{{- if $element.Is_basic }}
    {{ $element.Column_kind.Type}} {{$element.Column_name }};
{{ end }}
{{- end }}

    {{.Name}}_schema::TVaddr _var_ptrs[{{.Name}}_schema::VAR_POOL_NUM];
    {{.Name}}_schema::TPrimaryKey primary_key() const {
        return {{.Name}}_schema::TPrimaryKey({{.Primary_key.Name}});
    }
};

class {{.Name}}_object_tuple {
public:
    DEFINE_DATA_TUPLE({{.Name}}_object_tuple, {{.Name}}_schema)

${object_tuple_basic_fields}
${object_tuple_string_fields}
${object_tuple_vector_fields}
${object_tuple_binary_fields}
${object_tuple_portal_fields}

{{- range $index, $element := .Columns -}}

{{ if $element.Is_string}}
    DECLARE_DATA_TUPLE_REF_STRING_MEMBER({{$element.Column_name}}, {{$element.Length}});
{{ else if $element.Is_array}}
    DECLARE_DATA_TUPLE_REF_VECTOR_MEMBER({{$element.Column_kind.Type}}, {{$element.Column_name}}, {{$element.Length}});
{{ else if $element.Is_binary}}
    DECLARE_DATA_TUPLE_REF_BINARY_MEMBER({{$element.Column_name}}, {{$element.Length}});
{{ else }}
{{$element.Column_kind.Type}}, {{$element.Column_name}};
{{ end}}

{{- end -}}

public:
    void clear_portal_fields()
    {
${clear_object_tuple_portal_fields}
    }
    void init_fields_status() {
${init_fields_status}
    }
};

class {{.Name}}_object_tuple_ref {
public:
    DEFINE_DATA_TUPLE_REF({{.Name}}_object_tuple_ref, {{.Name}}_schema)

${object_tuple_ref_basic_fields}
DEFINE_DATA_TUPLE_REF_MEMBER

${object_tuple_ref_string_fields}
{{- range $index, $element := .Columns -}}

{{ if $element.Is_string}}
    DECLARE_DATA_TUPLE_REF_STRING_MEMBER({{$element.Column_name}}, {{$element.Length}});
{{ else if $element.Is_array}}
    DECLARE_DATA_TUPLE_REF_VECTOR_MEMBER({{$element.Column_kind.Type}}, {{$element.Column_name}}, {{$element.Length}});
{{ else if $element.Is_binary}}
    DECLARE_DATA_TUPLE_REF_BINARY_MEMBER({{$element.Column_name}}, {{$element.Length}});
{{ else if $element.Is_pb_expand}}
    DEFINE_DATA_TUPLE_PBEXPAND_REF_MEMBER({{$element.Column_kind.Type}}, {{$element.Column_name}});
{{ else }}
{{$element.Column_kind.Type}}, {{$element.Column_name}};
{{ end}}

{{- end -}}
};

class {{.Name}} : public galois::gtable::memory_data_table<{{.Name}}_schema> {
public:
    typedef {{.Name}}_schema TSchema;

    explicit {{.Name}}(const std::string &name)
        : galois::gtable::memory_data_table<{{.Name}}_schema>(name) {
        // nothing.
    }

    virtual ~{{.Name}}() {
        // nothing.
    }

    friend class {{.Name}}_object_tuple_ref;

public:
    int insert_var_pools(TBasicTuple &basic, const TObjectTuple &tuple);
    void get_var_field_name(char* name, uint32_t name_len, uint32_t idx) const {
{{- range $index, $element := .Columns}}
        if (idx == {{ $element.UpperName}} ) { 
            snprintf(name, name_len, "%s", "{{- $element.UpperName -}}"); 
            return; 
        }
{{- end}}
        snprintf(name, name_len, "%s", "ERROR");
    }

protected:
    int do_create_basic_pool(TBasicPool* p_basic_pool);

public:
    enum {
{{- range $index, $element := .Columns}}
        {{ $element.UpperName }} = {{ $index }},
{{- end}}
    };
};

template <typename TOstream>
inline TOstream &operator<<(TOstream &os, const {{.Name}}_object_tuple_ref &data) {
${data_table_dump_commands}
    return os;
}

}