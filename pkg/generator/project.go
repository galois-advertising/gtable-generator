//solopointer1202@gmail.com
package generator

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
	Queries       []Query
	ValueGetters  []ValueGetter
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
		if dsi, ok := _m[dv.DatasourceName]; ok {
			d.Datasources[dsi].Dataviews = append(d.Datasources[dsi].Dataviews,
				&d.Dataviews[dvi])
			log.Printf("Dataview:[%s] -> Datasource:[%s]", dv.Name, d.Datasources[dsi].Name)
		} else {
			panic(fmt.Sprintf("Cannot find datasource[%s] of dataview[%s]",
				dv.DatasourceName, dv.Name))
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
			log.Printf("link [%s] -> [%s]", d.Dataviews[idv].Name, du.Name)
		} else {
			log.Fatalf("dataupdator [%s]->[%s] : from not found", du.From, du.To)
		}
	}
	return nil
}

func (d *Project) build_datatable() error {
	_m := make(map[string]uint32)
	for idt, dt := range d.Datatables {
		_m[dt.Name] = uint32(idt)
		for icol, col := range d.Datatables[idt].Columns {
			d.Datatables[idt].Columns[icol].IsPrimarykey = false
			for _, pk := range d.Datatables[idt].Primary_key.Keys {
				if col.Column_name == pk {
					d.Datatables[idt].Columns[icol].IsPrimarykey = true
				}
			}
		}
	}
	for iiu, iu := range d.Indexupdators {
		if idt, ok := _m[iu.From]; ok {
			d.Datatables[idt].Indexupdators =
				append(d.Datatables[idt].Indexupdators, &d.Indexupdators[iiu])
			log.Printf("link [%s] -> [%s]", d.Datatables[idt].Name, iu.Name)
		} else {
			log.Fatalf("indexupdator [%s]->[%s] : from not found", iu.From, iu.To)
		}
	}
	return nil
}

func (d *Project) build_dataupdator() error {
	_m_dataview := make(map[string]uint32)
	for idx, dv := range d.Dataviews {
		_m_dataview[dv.Name] = uint32(idx)
	}
	_m_datatable := make(map[string]uint32)
	for idx, dt := range d.Datatables {
		_m_datatable[dt.Name] = uint32(idx)
	}
	for idx, du := range d.Dataupdators {
		log.Printf("Setting %s", du.Name)
		if ifrom, ok := _m_dataview[du.From]; ok {
			d.Dataupdators[idx].From_dataview = &d.Dataviews[ifrom]
		} else {
			msg := fmt.Sprintf("Cannot find the dataview of \"%s\":[%s]", du.Name, du.From)
			log.Fatal(msg)
			return errors.New(msg)
		}
		if ito, ok := _m_datatable[du.To]; ok {
			d.Dataupdators[idx].To_datatable = &d.Datatables[ito]
			for idtcol, dtcol := range d.Dataupdators[idx].To_datatable.Columns {
				column_from, err := d.Dataupdators[idx].From_dataview.IsDerivative(dtcol.Column_name)
				if err != nil {
					return err
				} else {
					d.Dataupdators[idx].To_datatable.Columns[idtcol].IsDerivative = column_from
				}
			}
		} else {
			msg := fmt.Sprintf("Cannot find the datatable of \"%s\":[%s]", du.Name, du.From)
			log.Fatal(msg)
			return errors.New(msg)
		}
	}

	return nil
}

func (d *Project) build_indexupdator() error {
	_m_datatable := make(map[string]uint32)
	for idx, dt := range d.Datatables {
		_m_datatable[dt.Name] = uint32(idx)
	}
	for idx, it := range d.Indexupdators {
		log.Printf("Setting %s", it.Name)
		if ifrom, ok := _m_datatable[it.From]; ok {
			d.Indexupdators[idx].From_datatable = &d.Datatables[ifrom]
		} else {
			msg := fmt.Sprintf("Cannot find the datatable of \"%s\":[%s]", it.Name, it.From)
			log.Fatal(msg)
			return errors.New(msg)
		}
	}

	return nil
}

func (d *Project) build_indextable() error {
	for iit, it := range d.Indextables {
		for _, dt := range d.Datatables {
			if it.OnTable == dt.Name {
				for _, col := range dt.Columns {
					if col.Column_name == it.OnColumn {
						if d.Indextables[iit].KeyType == "" ||
							col.Column_kind.Kind == "original" {
							if col.IsArray() {
								d.Indextables[iit].KeyType = ""
								msg := fmt.Sprintf("Array [%s] can not be used for index key",
									col.Column_name)
								log.Fatal(msg)
							} else if col.IsString() {
								d.Indextables[iit].KeyType = "std::string"
							} else {
								d.Indextables[iit].KeyType = col.Column_kind.Type
							}
						}
					}
				}
			}
		}
		if d.Indextables[iit].KeyType == "" {
			msg := fmt.Sprintf("Cannot find type of %s::%s", it.OnTable, it.OnColumn)
			log.Fatal(msg)
			return errors.New(msg)
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
	d.ValueGetters = []ValueGetter{}
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
				log.Printf("Loading [%s]", path)
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
	if err := d.build_dataupdator(); err != nil {
		errs := fmt.Sprintf("Build dataupdator fail for %s", err.Error())
		log.Fatal(errs)
		return errors.New(errs)
	}
	if err := d.build_datatable(); err != nil {
		errs := fmt.Sprintf("Build datatable fail for %s", err.Error())
		log.Fatal(errs)
		return errors.New(errs)
	}
	if err := d.build_indextable(); err != nil {
		errs := fmt.Sprintf("Build indextable fail for %s", err.Error())
		log.Fatal(errs)
		return errors.New(errs)
	}
	if err := d.build_indexupdator(); err != nil {
		errs := fmt.Sprintf("Build indexupdator fail for %s", err.Error())
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

func (d *Project) LoadGQL(gql_path string) error {
	if !DirExists(gql_path) {
		errs := fmt.Sprintf("%s not exists or not a directory.", gql_path)
		log.Fatal(errs)
		return errors.New(errs)
	}
	filepath.Walk(gql_path,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if !f.IsDir() && strings.HasSuffix(path, ".gql.xml") {
				log.Printf("Loading [%s]", path)
				r, err := ioutil.ReadFile(path)
				if err != nil {
					return err
				}
				var gql GqlDefines
				if err := xml.Unmarshal(r, &gql); err != nil {
					log.Fatal(err)
				} else {
					gql.Setup()
					for _, query := range gql.Queries {
						d.Queries = append(d.Queries, query)
					}
					// set ValueGetters
				}
			}
			return nil
		})
	return nil
}

func (d *Project) Generate(templates_path string, out_path string) error {
	if !DirExists(templates_path) {
		errs := fmt.Sprintf("Template path [%s] not exists or not a directory.", templates_path)
		log.Fatal(errs)
		return errors.New(errs)
	}
	if !DirExists(out_path) {
		errs := fmt.Sprintf("Output path [%s] not exists or not a directory.", out_path)
		log.Fatal(errs)
		return errors.New(errs)
	}
	var tmps GtableTemplates
	if err := tmps.Init(templates_path); err == nil {
		for _, ds := range d.Datasources {
			tmps.Generate(out_path, &ds)
		}
		for _, dv := range d.Dataviews {
			tmps.Generate(out_path, &dv)
		}
		for _, du := range d.Dataupdators {
			tmps.Generate(out_path, &du)
		}
		for _, iu := range d.Indexupdators {
			tmps.Generate(out_path, &iu)
		}
		for _, dt := range d.Datatables {
			tmps.Generate(out_path, &dt)
		}
		for _, it := range d.Indextables {
			tmps.Generate(out_path, &it)
		}
		for _, qy := range d.Queries {
			tmps.Generate(out_path, &qy)
		}
		for _, vg := range d.ValueGetters {
			tmps.Generate(out_path, &vg)
		}
	}
	tmps.generate_project(out_path, d)
	return nil
}
