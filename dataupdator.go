// solopointer1202@gmail.com
package main

type Dataupdator struct {
	Name      string   `xml:"name"`
	From      string   `xml:"from"`
	To        string   `xml:"to"`
	Udf       string   `xml:"properties>udf"`
	Type      string   `xml:"properties>type"`
	Notations []string `xml:"notations>notation"`
	Namespace string
	Handler   string
	Cppcode   string
}

func (d *Dataupdator) SetNamespace(_ns string) {
	d.Namespace = _ns
}

func (d *Dataupdator) SetHandler(_hd string) {
	d.Handler = _hd
}

func (d *Dataupdator) SetCppcode(_cpp string) {
	d.Cppcode = _cpp
}

func (d *Dataupdator) Setup() error {
	return nil
}
