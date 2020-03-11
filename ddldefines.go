//solopointer1202@gmail.com
package main

import "log"

type Iddl interface {
	SetNamespace(string)
	SetHandler(string)
	SetCppcode(string)
	Setup() error
}

type DdlDefines struct {
	Parser_build_time string         `xml:"parser_build_time"`
	Cppcode           string         `xml:"cppcode"`
	Handler           string         `xml:"handler"`
	Namespace         string         `xml:"namespace"`
	Datasources       []Datasource   `xml:"datasource"`
	Dataviews         []Dataview     `xml:"dataview"`
	Dataupdators      []Dataupdator  `xml:"dataupdator"`
	Datatables        []Datatable    `xml:"datatable"`
	Indexupdators     []Indexupdator `xml:"indexupdator"`
	Indextables       []Indextable   `xml:"indextable"`
}

func (d *DdlDefines) Setup() error {
	log.Printf("DDL namespace:%s", d.Namespace)
	for i, _ := range d.Datasources {
		d.Datasources[i].SetNamespace(d.Namespace)
		d.Datasources[i].SetHandler(d.Handler)
		d.Datasources[i].SetCppcode(d.Cppcode)
		d.Datasources[i].Setup()
	}
	for i, _ := range d.Dataviews {
		d.Dataviews[i].SetNamespace(d.Namespace)
		d.Dataviews[i].SetHandler(d.Handler)
		d.Dataviews[i].SetCppcode(d.Cppcode)
		d.Dataviews[i].Setup()
	}
	for i, _ := range d.Dataupdators {
		d.Dataupdators[i].SetNamespace(d.Namespace)
		d.Dataupdators[i].SetHandler(d.Handler)
		d.Dataupdators[i].SetCppcode(d.Cppcode)
		d.Dataupdators[i].Setup()
	}
	for i, _ := range d.Datatables {
		d.Datatables[i].SetNamespace(d.Namespace)
		d.Datatables[i].SetHandler(d.Handler)
		d.Datatables[i].SetCppcode(d.Cppcode)
		d.Datatables[i].Setup()
	}
	for i, _ := range d.Indexupdators {
		d.Indexupdators[i].SetNamespace(d.Namespace)
		d.Indexupdators[i].SetHandler(d.Handler)
		d.Indexupdators[i].SetCppcode(d.Cppcode)
		d.Indexupdators[i].Setup()
	}
	for i, _ := range d.Indextables {
		d.Indextables[i].SetNamespace(d.Namespace)
		d.Indextables[i].SetHandler(d.Handler)
		d.Indextables[i].SetCppcode(d.Cppcode)
		d.Indextables[i].Setup()
	}
	return nil
}
