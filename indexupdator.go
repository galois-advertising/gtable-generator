// solopointer1202@gmail.com
package main

import (
	"errors"
	"fmt"
	"strings"
)

type Indexupdator struct {
	Name       string     `xml:"name"`
	From       string     `xml:"from"`
	To         string     `xml:"to"`
	Properties []Property `xml:"property"`
	Notations  []string   `xml:"notations>notation"`
	Namespace  string
	Handler    string
	Cppcode    string
}

func (d *Indexupdator) Get_type() (string, error) {
	for _, v := range d.Properties {
		if v.Name == "type" {
			return v.Value, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Could not find type of %s", d.Name))
}

func (d *Indexupdator) Get_udf() (string, error) {
	for _, v := range d.Properties {
		if v.Name == "udf" {
			return v.Value, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Could not find udf of %s", d.Name))
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
	d.Name = strings.Replace(d.Name, "|", "_to_", -1)
	return nil
}
