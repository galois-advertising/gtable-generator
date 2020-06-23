package generator

type FieldCondition struct {
	Id     string   `xml:"id"`
	Op     string   `xml:"type"`
	Fields []string `xml:"field"`
	Left   Field
	Right  Field
}

func (fc *FieldCondition) Setup() (ok bool) {
	if fc.Op == "=" {
		fc.Op = "=="
	}

	return true
}
