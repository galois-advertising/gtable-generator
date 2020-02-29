// solopointer1202@gmail.com
package main

type Indexupdator struct {
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

func (d *Indexupdator) SetNamespace(_ns string) {
	d.Namespace = _ns
}

func (d *Indexupdator) SetHandler(_hd string) {
	d.Handler = _hd
}

func (d *Indexupdator) SetCppcode(_cpp string) {
	d.Cppcode = _cpp
}

func (d *Indexupdator) Setup() error {
	return nil
}
