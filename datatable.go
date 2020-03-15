// solopointer1202@gmail.com
package main

import (
	"fmt"
	"log"
	"strings"
)

type Datatable struct {
	Name        string     `xml:"name"`
	Columns     []Column   `xml:"columns_node>column_node"`
	Primary_key Primarykey `xml:"primary_key"`
	Notations   []string   `xml:"notations>notation"`
	Namespace   string
	Handler     string
	Cppcode     string
}

type Primarykey struct {
	Text string `xml:",chardata"`
	Type string `xml:"type,attr"`
	Keys []string
}

func (d *Datatable) Is_primarykey(name string) bool {
	for _, col := range d.Columns {
		if col.Column_name == name {
			for _, pk := range d.Primary_key.Keys {
				if name == pk {
					return true
				}
			}
			return false
		}
	}
	msg := fmt.Sprintf("Is_primarykey: Cannot find column [%s] in [%s]", name, d.Name)
	log.Fatal(msg)
	return false
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
	d.Primary_key.Keys = strings.Split(d.Primary_key.Text, ",")
	return nil
}
