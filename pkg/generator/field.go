package generator

type Field struct {
	Name      string `xml:",chardata"`
	Apply     string `xml:"apply,attr"`
	TableType string
	FieldName string
	FieldType string
	ParamType string
}

func (f *Field) Setup() (ok bool) {
	//func_type = tag_text(xml_node, "func_type")
	//if (func_type != ""):
	//    self.set_is_func_field()
	//    self.func_type = func_type
	//    self.name = tag_text(xml_node, "name")
	//else:
	//    self.name = node_text(xml_node, "field")

	//self.ori_name = self.name
	//if self.name.startswith('$'):
	//    self.set_placeholde_info(self.name)
	//    self.field_name = self.name
	//else:
	//    self.check_schema(schema_handle)
	return true
}
