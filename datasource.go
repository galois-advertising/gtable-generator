// solopointer1202@gmail.com
package main

type Datasource struct {
	Name      string   `xml:"name"`
	Type      string   `xml:"properties>type"`
	Protourl  string   `xml:"properties>protourl"`
	Notations []string `xml:"notations>notation"`
	Dataviews []*Dataview
	Namespace string
	Handler   string
	Cppcode   string
}

func (d *Datasource) SetNamespace(_ns string) {
	d.Namespace = _ns
}

func (d *Datasource) SetHandler(_hd string) {
	d.Handler = _hd
}

func (d *Datasource) SetCppcode(_cpp string) {
	d.Cppcode = _cpp
}

func (d *Datasource) Setup() error {
	d.Dataviews = []*Dataview{}
	return nil
}
