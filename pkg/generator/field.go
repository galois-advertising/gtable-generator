package generator

import (
	"errors"
	"fmt"
	"strings"
)

type Field struct {
	Name  string `xml:",chardata"`
	Apply string `xml:"apply,attr"`
}

func (f *Field) GetName() (string, bool) {
	if table, tok := f.TableName(); tok == nil {
		if field, fok := f.FieldName(); fok == nil {
			return table + "_" + field, true
		}
	}
	return "", false
}

func (f *Field) IsPlaceHolder() bool {
	return f.Name[0] == '$'
}

func (f *Field) PlaceHolderName() (string, error) {
	if len(f.Name) < 2 {
		return "", errors.New(fmt.Sprintf("[gtable] Invalid placeholder name [%s]", f.Name))
	}
	return f.Name[1:], nil
}

func (f *Field) IsTableField() bool {
	for i := 0; i < len(f.Name); i++ {
		if f.Name[i] == '.' && i != 0 && i != len(f.Name)-1 {
			return true
		}
	}
	return false
}

func (f *Field) TableName() (string, error) {
	if !f.IsPlaceHolder() {
		spls := strings.Split(f.Name, ".")
		if len(spls) == 2 && len(spls[0]) > 0 {
			return spls[0], nil
		}
	}
	return "", errors.New(fmt.Sprintf("[gtable] Invalid TableName [%s]", f.Name))
}

func (f *Field) FieldName() (string, error) {
	if f.IsPlaceHolder() {
		if len(f.Name) > 1 {
			return f.Name[1:], nil
		}
	} else {
		spls := strings.Split(f.Name, ".")
		if len(spls) == 2 && len(spls[1]) > 0 {
			return spls[1], nil
		}
	}
	return "", errors.New(fmt.Sprintf("[gtable] Invalid TableName [%s]", f.Name))
}

func (f *Field) ApplyFunc() (string, bool) {
	return f.Apply, len(f.Apply) != 0
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
