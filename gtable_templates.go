package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

type GtableTemplates struct {
	tmpls map[string]*template.Template
}

func (gt *GtableTemplates) Init(root string) error {
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

func (gt *GtableTemplates) generate_dataview(out_path string, dv *Dataview) {
	if udf, err := dv.GetUDF(); err == nil {
		h_file, err := os.Create(fmt.Sprintf("%s/include/%s.h.example", out_path, udf))
		if err != nil {
			log.Fatal(err)
		}
		defer h_file.Close()
		tmpl := gt.tmpls["dataview_udf_example.t"]
		if err := tmpl.Execute(h_file, dv); err != nil {
			panic(err.Error())
		}
	}
	h_file, err := os.Create(fmt.Sprintf("%s/include/%s.h", out_path, dv.Name))
	if err != nil {
		log.Fatal(err)
	}
	defer h_file.Close()
	tmpl := gt.tmpls["dataview_h.t"]
	if err := tmpl.Execute(h_file, dv); err != nil {
		panic(err.Error())
	}
}

func (gt *GtableTemplates) generate_datatable(out_path string, dt *Datatable) {
	log.Printf("Processing %s", dt.Name)
	h_file, err := os.Create(fmt.Sprintf("%s/include/%s.h", out_path, dt.Name))
	if err != nil {
		log.Fatal(err)
	}
	defer h_file.Close()
	tmpl := gt.tmpls["datatable_h.t"]
	if err := tmpl.Execute(h_file, dt); err != nil {
		panic(err.Error())
	}
	cpp_file, err := os.Create(fmt.Sprintf("%s/src/%s.cpp", out_path, dt.Name))
	if err != nil {
		log.Fatal(err)
	}
	defer cpp_file.Close()
	tmpl = gt.tmpls["datatable_cpp.t"]
	if err := tmpl.Execute(cpp_file, dt); err != nil {
		panic(err.Error())
	}
}

func (gt *GtableTemplates) generate_datasource_databus(out_path string, ds *Datasource) {
	log.Printf("Processing %s", ds.Name)
	h_file, err := os.Create(fmt.Sprintf("%s/include/%s.h", out_path, ds.Name))
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
	cpp_file, err := os.Create(fmt.Sprintf("%s/src/%s.cpp", out_path, ds.Name))
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

func (gt *GtableTemplates) generate_datasource(out_path string, ds *Datasource) {
	dtype, err := ds.Get_type()
	log.Printf(dtype)
	if err == nil {
		switch dtype {
		case "databus":
			gt.generate_datasource_databus(out_path, ds)
		default:
		}
	}
}

func (gt *GtableTemplates) generate_dataupdator(out_path string, du *Dataupdator) {
	log.Printf("Processing %s", du.Name)
	h_file, err := os.Create(fmt.Sprintf("%s/include/%s.h", out_path, du.Name))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer h_file.Close()
	if tmpl, ok := gt.tmpls["dataupdator_h.t"]; ok {
		if err := tmpl.Execute(h_file, du); err != nil {
			panic(err.Error())
		}
	} else {
		panic(fmt.Sprintf("Cannot find .h template for dataupdator"))
	}
}

func (gt *GtableTemplates) generate_indextable(out_path string, it *Indextable) {
	log.Printf("Processing %s", it.Name)
	h_file, err := os.Create(fmt.Sprintf("%s/include/%s.h", out_path, it.Name))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer h_file.Close()
	if tmpl, ok := gt.tmpls["indextable_h.t"]; ok {
		if err := tmpl.Execute(h_file, it); err != nil {
			panic(err.Error())
		}
	} else {
		panic(fmt.Sprintf("Cannot find .h template for indextable"))
	}
}

func (gt *GtableTemplates) generate_project(out_path string, p *Project) error {
	h_file, err := os.Create(fmt.Sprintf("%s/include/project.h", out_path))
	if err != nil {
		msg := fmt.Sprintf("Create file [%s] fail for %s", "env.h", err.Error())
		log.Fatal(msg)
		return errors.New(msg)
	}
	defer h_file.Close()
	tmpl := gt.tmpls["project_h.t"]
	if err := tmpl.Execute(h_file, p); err != nil {
		panic(err.Error())
	}
	return nil
}

func (gt *GtableTemplates) Generate(out_path string, data interface{}) {
	if !DirExists(out_path) {
		panic(fmt.Sprintf("Output path [%s] does not exists.", out_path))
	}
	switch data.(type) {
	case *Datatable:
		gt.generate_datatable(out_path, data.(*Datatable))
	case *Dataview:
		gt.generate_dataview(out_path, data.(*Dataview))
	case *Datasource:
		gt.generate_datasource(out_path, data.(*Datasource))
	case *Dataupdator:
		gt.generate_dataupdator(out_path, data.(*Dataupdator))
	case *Indextable:
		gt.generate_indextable(out_path, data.(*Indextable))
	case *Project:
		gt.generate_project(out_path, data.(*Project))
	default:
		panic(fmt.Sprintf("Unknow type"))
	}
}
