package generator

type LogicCondition struct {
	Id    string   `xml:"id"`
	Op    string   `xml:"type"`
	SubId []string `xml:"sub_conditioner"`
}
