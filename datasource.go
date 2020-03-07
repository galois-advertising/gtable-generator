// solopointer1202@gmail.com
package main

import (
	"errors"
	"fmt"
)

type Datasource struct {
	Name       string     `xml:"name"`
	Properties []Property `xml:"property"`
	Notations  []string   `xml:"notations>notation"`
	Dataviews  []*Dataview
	Namespace  string
	Handler    string
	Cppcode    string
}

func (d *Datasource) Get_protourl() (string, error) {
	for _, v := range d.Properties {
		if v.Name == "protourl" {
			return v.Value, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Could not find protourl of %s", d.Name))
}

func (d *Datasource) Get_type() (string, error) {
	for _, v := range d.Properties {
		if v.Name == "type" {
			return v.Value, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Could not find type of %s", d.Name))
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
