package generator

type Query struct {
	Name            string           `xml:"name"`
	Columns         []SelectColumns  `xml:"select>column"`
	Tables          []JoinedTable    `xml:"from>table"`
	FieldConditions []FieldCondition `xml:"where>field_conditioner"`
	LogicConditions []LogicCondition `xml:"where>logic_conditioner"`
	Namespace       string
	FieldsMap       map[string]*Field
}

func (ql *Query) setup() (ok bool) {
	ql.FieldsMap = map[string]*Field{}
	if !ql.collectValueGetters() {
		return false
	}
	return true
}

func (ql *Query) collectValueGetters() bool {
	for i, tab := range ql.Tables {
		if tab.ScanLimit.Name != "" {
			ql.FieldsMap[tab.ScanLimit.Name] = &ql.Tables[i].ScanLimit
		}
		if tab.EachScanLimit.Name != "" {
			ql.FieldsMap[tab.EachScanLimit.Name] = &ql.Tables[i].EachScanLimit
		}
		if tab.ResultLimit.Name != "" {
			ql.FieldsMap[tab.ResultLimit.Name] = &ql.Tables[i].ResultLimit
		}
		if tab.EachResultLimit.Name != "" {
			ql.FieldsMap[tab.EachResultLimit.Name] = &ql.Tables[i].EachResultLimit
		}
		for fi, field := range tab.LeftOnColumns {
			ql.FieldsMap[field.Name] = &ql.Tables[i].LeftOnColumns[fi]
		}
		for fi, field := range tab.RightOnColumns {
			ql.FieldsMap[field.Name] = &ql.Tables[i].RightOnColumns[fi]
		}
	}
	return true
}

func (ql *Query) regField(f *Field) {
	if _, has := f.ApplyFunc(); !has {
		if _, ok := ql.FieldsMap[f.Name]; !ok {
			ql.FieldsMap[f.Name] = f
		}
	}
}
