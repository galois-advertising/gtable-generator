// solopointer1202@gmail.com
package generator

import (
	"errors"
	"fmt"
	"strings"
)

type Dataupdator struct {
	Name          string     `xml:"name"`
	From          string     `xml:"from"`
	To            string     `xml:"to"`
	Properties    []Property `xml:"property"`
	Notations     []string   `xml:"notations>notation"`
	From_dataview *Dataview
	To_datatable  *Datatable
	Namespace     string
	Handler       string
	Cppcode       string
}

func (d *Dataupdator) Get_type() (string, error) {
	for _, v := range d.Properties {
		if v.Name == "type" {
			return v.Value, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Could not find type of %s", d.Name))
}

func (d *Dataupdator) GetUDF() (string, error) {
	for _, v := range d.Properties {
		if v.Name == "udf" {
			return v.Value, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Could not find udf of %s", d.Name))
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
	d.Name = strings.Replace(d.Name, "|", "_to_", -1)
	return nil
}
