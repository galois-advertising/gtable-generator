package generator 

type Property struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}
