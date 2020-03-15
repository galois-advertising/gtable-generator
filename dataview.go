// solopointer1202@gmail.com
package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

type Dataview struct {
	Name              string     `xml:"name"`
	Channel           string     `xml:"channel"`
	Columns           []Column   `xml:"columns_node>column_node"`
	Notations         []string   `xml:"notations>notation"`
	Properties        []Property `xml:"property"`
	Namespace         string
	Handler           string
	Cppcode           string
	DatasourceName    string
	DatasourceChannel string
	Dataupdators      []*Dataupdator
}

func (d *Dataview) IsDerivative(name string) (string, error) {
	result := ""
	for _, col := range d.Columns {
		if col.Column_name == name {
			if len(result) == 0 || result == "original" {
				result = col.IsDerivative
			}
		}
	}
	if len(result) == 0 {
		msg := fmt.Sprintf("IsDerivative: Cannot find column [%s] in [%s]", name, d.Name)
		log.Fatal(msg)
		return "", errors.New(msg)
	} else {
		return result, nil
	}
}

func (d *Dataview) HasUDF() bool {
	for _, v := range d.Properties {
		if v.Name == "udf" {
			return true
		}
	}
	return false
}

func (d *Dataview) GetUDF() (string, error) {
	for _, v := range d.Properties {
		if v.Name == "udf" {
			return v.Value, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Could not find udf of %s", d.Name))
}

func (d *Dataview) SetNamespace(_ns string) {
	d.Namespace = _ns
}

func (d *Dataview) SetHandler(_hd string) {
	d.Handler = _hd
}

func (d *Dataview) SetCppcode(_cpp string) {
	d.Cppcode = _cpp
}

func (d *Dataview) Setup() error {
	d.Dataupdators = make([]*Dataupdator, 0, 0)
	spl := strings.Split(d.Channel, "::")
	if len(spl) == 2 {
		d.DatasourceName = spl[0]
		d.DatasourceChannel = spl[1]
		log.Printf("Dataview:[%s] on [%s]::[%s]", d.Name,
			d.DatasourceName, d.DatasourceChannel)
	} else {
		panic(fmt.Sprintf("%s is not a valid datasource channel", d.Channel))
	}
	return nil
}
