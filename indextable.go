// solopointer1202@gmail.com
package main

type Indextable struct {
	Name                     string   `xml:"name"`
	Ontable                  string   `xml:"on_table"`
	OnColumn                 string   `xml:"on_column"`
	Columns                  []Column `xml:"columns_node>column_node"`
	Notations                []string `xml:"notations>notation"`
	Include_dataview_headers []string
	Namespace                string
	Handler                  string
	Cppcode                  string
}

func (d *Indextable) SetNamespace(_ns string) {
	d.Namespace = _ns
}

func (d *Indextable) SetHandler(_hd string) {
	d.Handler = _hd
}

func (d *Indextable) SetCppcode(_cpp string) {
	d.Cppcode = _cpp
}

func (d *Indextable) Setup() error {
	d.Include_dataview_headers = []string{"common.h", "b.h"}
	return nil
}
