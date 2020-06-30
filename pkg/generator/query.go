package generator

import "fmt"

type Query struct {
	Name            string           `xml:"name"`
	Columns         []SelectColumns  `xml:"select>column"`
	Tables          []JoinedTable    `xml:"from>table"`
	FieldConditions []FieldCondition `xml:"where>field_conditioner"`
	LogicConditions []LogicCondition `xml:"where>logic_conditioner"`
	Namespace       string
	fieldsMap       map[string]*Field
}

func (ql *Query) setup() (ok bool) {
	ql.fieldsMap = map[string]*Field{}
	if !ql.collectValueGetters() {
		return false
	}
	return true
}

func (ql *Query) collectValueGetters() bool {
	for i, tab := range ql.Tables {
		if tab.ScanLimit.Name != "" {
			ql.fieldsMap[tab.ScanLimit.Name] = &ql.Tables[i].ScanLimit
		}
		if tab.EachScanLimit.Name != "" {
			ql.fieldsMap[tab.EachScanLimit.Name] = &ql.Tables[i].EachScanLimit
		}
		if tab.ResultLimit.Name != "" {
			ql.fieldsMap[tab.ResultLimit.Name] = &ql.Tables[i].ResultLimit
		}
		if tab.EachResultLimit.Name != "" {
			ql.fieldsMap[tab.EachResultLimit.Name] = &ql.Tables[i].EachResultLimit
		}
		for fi, field := range tab.LeftOnColumns {
			ql.fieldsMap[field.Name] = &ql.Tables[i].LeftOnColumns[fi]
		}
		for fi, field := range tab.RightOnColumns {
			ql.fieldsMap[field.Name] = &ql.Tables[i].RightOnColumns[fi]
		}
	}
	for k, v := range ql.fieldsMap {
		fmt.Println(k, "->", v)
	}
	return true
}

func (ql *Query) regField(f *Field) {
	if _, has := f.ApplyFunc(); !has {
		if _, ok := ql.fieldsMap[f.Name]; !ok {
			ql.fieldsMap[f.Name] = f
		}
	}
}

func (ql *Query) SetupValueGetter() (ok bool) {
	for i := 0; i < len(ql.FieldConditions); i++ {

	}
	return true

}
