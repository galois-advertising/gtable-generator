package generator

type FieldCondition struct {
	Id     string  `xml:"id"`
	Op     string  `xml:"type"`
	Fields []Field `xml:"field"`
}

func (fc *FieldCondition) Setup() (ok bool) {
	if fc.Op == "=" {
		fc.Op = "=="
	}

	return true
}
