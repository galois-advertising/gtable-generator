// solopointer1202@gmail.com
package main

import (
	"fmt"
	"log"
	"strings"
)

type Dataview struct {
	Name               string   `xml:"name"`
	Channel            string   `xml:"channel"`
	Columns            []Column `xml:"columns_node>column_node"`
	Notations          []string `xml:"notations>notation"`
	Namespace          string
	Handler            string
	Cppcode            string
	Datasource_name    string
	Datasource_channel string
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
	spl := strings.Split(d.Channel, "::")
	if len(spl) == 2 {
		d.Datasource_name = spl[0]
		d.Datasource_channel = spl[1]
		log.Printf("Dataview:[%s] on [%s]::[%s]", d.Name,
			d.Datasource_name, d.Datasource_channel)
	} else {
		panic(fmt.Sprintf("%s is not a valid datasource channel", d.Channel))
	}
	return nil
}
