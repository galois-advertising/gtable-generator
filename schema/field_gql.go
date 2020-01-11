package gqlgenerator

import (
	"string"
)

type field struct {
    name ori_name string
    table_name table_type string
    field_name field_type string
    is_in_index is_pb_expand is_placeholder is_func_field bool
    func_type string
    schema_field string 
    param_type string 
}

func (f *field) load(xml string) {

}