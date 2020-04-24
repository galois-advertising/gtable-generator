package generator

type SelectColumns struct {
	Name string `xml:",chardata"`
}

type JoinedTable struct {
	Name            string   `xml:"name"`
	OnColumn        []string `xml:"on_column"`
	JoinType        string   `xml:"join_type,attr"`
	ScanLimit       string   `xml:"scan_limit,attr"`
	EachScanLimit   string   `xml:"each_scan_limit,attr"`
	ResultLimit     string   `xml:"result_limit,attr"`
	EachResultLimit string   `xml:"each_result_limit,attr"`
}

type FieldCondition struct {
	Id    string   `xml:"id"`
	Op    string   `xml:"type"`
	Field []string `xml:"field"`
}

type LogicCondition struct {
	Id    string   `xml:"id"`
	Op    string   `xml:"type"`
	SubId []string `xml:"sub_conditioner"`
}

type Query struct {
	Name            string           `xml:"name"`
	Columns         []SelectColumns  `xml:"select>column"`
	Tables          []JoinedTable    `xml:"from>table"`
	FieldConditions []FieldCondition `xml:"where>field_conditioner"`
	LogicConditions []LogicCondition `xml:"where>logic_conditioner"`
}

type GqlDefines struct {
	Queries []Query `xml:"gql>query"`
}

func (gd *GqlDefines) Setup() (ok bool) {
	return true
}
