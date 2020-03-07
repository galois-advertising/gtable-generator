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

func (gt *Gtable_templates) generate_dataview(out_path string, dv *Dataview) {
	if udf, err := dv.Get_udf(); err == nil {
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
}

func (gt *Gtable_templates) generate_datatable(out_path string, dt *Datatable) {
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

func (gt *Gtable_templates) generate_datasource_databus(out_path string, ds *Datasource) {
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

func (gt *Gtable_templates) generate_datasource(out_path string, ds *Datasource) {
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

func (gt *Gtable_templates) Generate(out_path string, data interface{}) {
	if !exists(out_path) {
		panic(fmt.Sprintf("[%s] does not exists.", out_path))
	}
	log.Printf("Start to generate")
	switch data.(type) {
	case *Datatable:
		gt.generate_datatable(out_path, data.(*Datatable))
	case *Dataview:
		gt.generate_dataview(out_path, data.(*Dataview))
	case *Datasource:
		gt.generate_datasource(out_path, data.(*Datasource))
	default:
		panic(fmt.Sprintf("Unknow type"))
	}
}
