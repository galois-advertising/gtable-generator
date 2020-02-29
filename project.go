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
	Indextalbes   []Indextable
}

func (d *Project) Stat() {
	log.Printf("Datasource\t\tlen:%d\n", len(d.Datasources))
	log.Printf("Dataviews\t\tlen:%d\n", len(d.Dataviews))
	log.Printf("Dataupdators\t\tlen:%d\n", len(d.Dataupdators))
	log.Printf("Datatables\t\tlen:%d\n", len(d.Datatables))
	log.Printf("Indexupdators\t\tlen:%d\n", len(d.Indexupdators))
	log.Printf("Indextalbes\t\tlen:%d\n", len(d.Indextalbes))
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
	return nil
}

func (d *Project) Init() error {
	d.Datasources = []Datasource{}
	d.Dataviews = []Dataview{}
	d.Dataupdators = []Dataupdator{}
	d.Datatables = []Datatable{}
	d.Indexupdators = []Indexupdator{}
	d.Indextalbes = []Indextable{}
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
					for _, dv := range ddl.Indextalbes {
						d.Indextalbes = append(d.Indextalbes, dv)
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
			tmps.Generate(out_path, ds)
		}
	}
	return nil
}
