package generator

type Query struct {
	Name            string           `xml:"name"`
	Columns         []SelectColumns  `xml:"select>column"`
	Tables          []JoinedTable    `xml:"from>table"`
	FieldConditions []FieldCondition `xml:"where>field_conditioner"`
	LogicConditions []LogicCondition `xml:"where>logic_conditioner"`
	Namespace       string
}

func (ql *Query) setup() (ok bool) {
	return true
}

func (ql *Query) SetupValueGetter() (ok bool) {
	for i := 0; i < len(ql.FieldConditions); i++ {

	}
	return true

}
