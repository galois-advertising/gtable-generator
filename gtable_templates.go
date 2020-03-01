package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

type Gtable_templates struct {
	tmpls map[string]*template.Template
}

func (gt *Gtable_templates) Init(root string) error {
	gt.tmpls = make(map[string]*template.Template)
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if len(path) > 3 && path[len(path)-2:] == ".t" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		name := path.Base(file)
		log.Printf("Loading template: %s:\t[%s]\n", name, file)
		t := template.Must(template.New(name).ParseFiles(file))
		gt.tmpls[name] = t
	}
	return nil

}

func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (gt *Gtable_templates) generate_datatable(out_path string, data interface{}) {
	dt := data.(Datatable)
	h_file, err := os.Create(fmt.Sprintf("%s/%s.h", out_path, dt.Name))
	if err != nil {
		log.Fatal(err)
	}
	defer h_file.Close()
	tmpl := gt.tmpls["datatable_h.t"]
	if err := tmpl.Execute(h_file, data); err != nil {
		panic(err.Error())
	}
	cpp_file, err := os.Create(fmt.Sprintf("%s/%s.cpp", out_path, dt.Name))
	if err != nil {
		log.Fatal(err)
	}
	defer cpp_file.Close()
	tmpl = gt.tmpls["datatable_cpp.t"]
	if err := tmpl.Execute(cpp_file, data); err != nil {
		panic(err.Error())
	}
}

func (gt *Gtable_templates) generate_datasource_databus(out_path string, ds *Datasource) {
	h_file, err := os.Create(fmt.Sprintf("%s/%s.h", out_path, ds.Name))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer h_file.Close()
	if tmpl, ok := gt.tmpls["databus_h1.t"]; ok {
		if err := tmpl.Execute(h_file, ds); err != nil {
			panic(err.Error())
		}
	} else {
		panic(fmt.Sprintf("Cannot find .h template for Datasource::databus"))
	}
	cpp_file, err := os.Create(fmt.Sprintf("%s/%s.cpp", out_path, ds.Name))
	if err != nil {
		log.Fatal(err)
	}
	defer h_file.Close()
	if tmpl, ok := gt.tmpls["databus_cpp.t"]; ok {
		if err := tmpl.Execute(cpp_file, ds); err != nil {
			panic(err.Error())
		}
	} else {
		panic(fmt.Sprintf("Cannot find .cpp template for Datasource::databus"))
	}
}

func (gt *Gtable_templates) generate_datasource(out_path string, data interface{}) {
	ds := data.(Datasource)
	switch ds.Type {
	case "databus":
		gt.generate_datasource_databus(out_path, &ds)
	default:
	}
}

func (gt *Gtable_templates) Generate(out_path string, data interface{}) {
	if !exists(out_path) {
		panic(fmt.Sprintf("[%s] does not exists.", out_path))
	}
	switch data.(type) {
	case Datatable:
		gt.generate_datatable(out_path, data)
	case Datasource:
		gt.generate_datasource(out_path, data)
	default:
		panic(fmt.Sprintf("Unknow type"))
	}
}