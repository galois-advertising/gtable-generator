//solopointer1202@gmail.com
package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Project struct {
	Datasources   []Datasource
	Dataviews     []Dataview
	Dataupdators  []Dataupdator
	Datatables    []Datatable
	Indexupdators []Indexupdator
	Indextables   []Indextable
	Namespace     string
	Handler       string
}

func (d *Project) Stat() {
	log.Printf("Datasource\t\tlen:%d\n", len(d.Datasources))
	log.Printf("Dataviews\t\tlen:%d\n", len(d.Dataviews))
	log.Printf("Dataupdators\t\tlen:%d\n", len(d.Dataupdators))
	log.Printf("Datatables\t\tlen:%d\n", len(d.Datatables))
	log.Printf("Indexupdators\t\tlen:%d\n", len(d.Indexupdators))
	log.Printf("Indextables\t\tlen:%d\n", len(d.Indextables))
}

func (d *Project) build_datasource() error {
	_m := make(map[string]uint32)
	for idx, ds := range d.Datasources {
		_m[ds.Name] = uint32(idx)
		log.Printf("Datasource:[%s]", ds.Name)
	}
	for dvi, dv := range d.Dataviews {
		if dsi, ok := _m[dv.Datasource_name]; ok {
			d.Datasources[dsi].Dataviews = append(d.Datasources[dsi].Dataviews,
				&d.Dataviews[dvi])
			log.Printf("Dataview:[%s] -> Datasource:[%s]", dv.Name, d.Datasources[dsi].Name)
		} else {
			panic(fmt.Sprintf("Cannot find datasource[%s] of dataview[%s]",
				dv.Datasource_name, dv.Name))
		}
	}
	return nil
}

func (d *Project) build_dataview() error {
	_m := make(map[string]uint32)
	for idx, dv := range d.Dataviews {
		_m[dv.Name] = uint32(idx)
		log.Printf("Dataview:[%s]", dv.Name)
	}
	for idu, du := range d.Dataupdators {
		if idv, ok := _m[du.From]; ok {
			d.Dataviews[idv].Dataupdators =
				append(d.Dataviews[idv].Dataupdators, &d.Dataupdators[idu])
			log.Printf("link [%s] -> [%s]", du.Name, d.Dataviews[idv].Name)
		} else {
			log.Fatalf("dataupdator [%s]->[%s] : from not found", du.From, du.To)
		}
	}
	return nil
}

func (d *Project) Init() {
	d.Datasources = []Datasource{}
	d.Dataviews = []Dataview{}
	d.Dataupdators = []Dataupdator{}
	d.Datatables = []Datatable{}
	d.Indexupdators = []Indexupdator{}
	d.Indextables = []Indextable{}
}

func (d *Project) check_and_set_handler() error {
	handler := make(map[string]bool)
	for _, v := range d.Datasources {
		handler[v.Handler] = true
	}
	for _, v := range d.Dataviews {
		handler[v.Handler] = true
	}
	for _, v := range d.Datatables {
		handler[v.Handler] = true
	}
	for _, v := range d.Dataupdators {
		handler[v.Handler] = true
	}
	for _, v := range d.Indextables {
		handler[v.Handler] = true
	}
	if len(handler) != 1 {
		errs := fmt.Sprintf("Handler not unique:%v", handler)
		log.Fatal(errs)
		return errors.New(errs)
	} else {
		log.Printf("Handler is unique:%v", handler)
		for k, _ := range handler {
			d.Handler = k
		}
	}
	return nil
}

func (d *Project) check_and_set_namespace() error {
	namespace := make(map[string]bool)
	for _, v := range d.Datasources {
		namespace[v.Namespace] = true
	}
	for _, v := range d.Dataviews {
		namespace[v.Namespace] = true
	}
	for _, v := range d.Datatables {
		namespace[v.Namespace] = true
	}
	for _, v := range d.Dataupdators {
		namespace[v.Namespace] = true
	}
	for _, v := range d.Indextables {
		namespace[v.Namespace] = true
	}
	if len(namespace) != 1 {
		errs := fmt.Sprintf("Namespace not unique:%v", namespace)
		log.Fatal(errs)
		return errors.New(errs)
	} else {
		log.Printf("Namespace is unique:%v", namespace)
		for k, _ := range namespace {
			d.Namespace = k
		}
	}
	return nil
}

func (d *Project) LoadDDL(ddl_path string) error {
	if !DirExists(ddl_path) {
		errs := fmt.Sprintf("%s not exists or not a directory.", ddl_path)
		log.Fatal(errs)
		return errors.New(errs)
	}
	filepath.Walk(ddl_path,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if !f.IsDir() && strings.HasSuffix(path, ".ddl.xml") {
				log.Printf("Processing: %s", path)
				r, err := ioutil.ReadFile(path)
				if err != nil {
					return err
				}
				var ddl DdlDefines
				if err := xml.Unmarshal(r, &ddl); err != nil {
					log.Fatal(err)
				} else {
					ddl.Setup()
					for _, ds := range ddl.Datasources {
						d.Datasources = append(d.Datasources, ds)
					}
					for _, dv := range ddl.Dataviews {
						d.Dataviews = append(d.Dataviews, dv)
					}
					for _, dv := range ddl.Dataupdators {
						d.Dataupdators = append(d.Dataupdators, dv)
					}
					for _, dv := range ddl.Datatables {
						d.Datatables = append(d.Datatables, dv)
					}
					for _, dv := range ddl.Indexupdators {
						d.Indexupdators = append(d.Indexupdators, dv)
					}
					for _, dv := range ddl.Indextables {
						d.Indextables = append(d.Indextables, dv)
					}
				}
			}
			return nil
		})
	if err := d.build_datasource(); err != nil {
		errs := fmt.Sprintf("Build datasource fail for %s", err.Error())
		log.Fatal(errs)
		return errors.New(errs)
	}
	if err := d.build_dataview(); err != nil {
		errs := fmt.Sprintf("Build dataview fail for %s", err.Error())
		log.Fatal(errs)
		return errors.New(errs)
	}
	if err := d.check_and_set_namespace(); err != nil {
		return err
	}
	if err := d.check_and_set_handler(); err != nil {
		return err
	}

	return nil
}

func (d *Project) Generate(templates_path string, out_path string) error {
	if !DirExists(templates_path) {
		errs := fmt.Sprintf("%s not exists or not a directory.", templates_path)
		log.Fatal(errs)
		return errors.New(errs)
	}
	if !DirExists(out_path) {
		errs := fmt.Sprintf("%s not exists or not a directory.", out_path)
		log.Fatal(errs)
		return errors.New(errs)
	}
	var tmps Gtable_templates
	if err := tmps.Init(templates_path); err == nil {
		for _, ds := range d.Datasources {
			tmps.Generate(out_path, &ds)
		}
		for _, dv := range d.Dataviews {
			tmps.Generate(out_path, &dv)
		}
	}
	tmps.generate_project(out_path, d)
	return nil
}
