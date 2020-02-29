// solopointer1202@gmail.com
package main

type Datatable struct {
	Name                     string     `xml:"name"`
	Columns                  []Column   `xml:"columns_node>column_node"`
	Primary_key              Primarykey `xml:"primary_key"`
	Notations                []string   `xml:"notations>notation"`
	Include_dataview_headers []string
	Namespace                string
	Handler                  string
	Cppcode                  string
}

type Primarykey struct {
	Name string `xml:",chardata"`
	Type string `xml:"type,attr"`
}

func (d *Datatable) SetNamespace(_ns string) {
	d.Namespace = _ns
}

func (d *Datatable) SetHandler(_hd string) {
	d.Handler = _hd
}

func (d *Datatable) SetCppcode(_cpp string) {
	d.Cppcode = _cpp
}

func (d *Datatable) Setup() error {
	d.Include_dataview_headers = []string{"common.h", "b.h"}
	d.Namespace = "galois::gtable"
	return nil
}
