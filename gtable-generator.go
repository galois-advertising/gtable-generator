package main 

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"github.com/galois-advertising/gtable-generator/pkg/generator"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:  %s [-t templates_dir] [-i ddl_path] [-o path]\nOptions:\n",
		os.Args[0])
	flag.PrintDefaults()
}

var (
	help           bool
	template_path  string
	ddl_input_path string
	gql_input_path string
	output_path    string
)

func init() {
	var default_template_path string
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	default_template_path = path.Join(filepath.Dir(ex), 
	    "../src/github.com/galois-advertising/gtable-generator/gtable-generator-templates")
	flag.BoolVar(&help, "h", false, "Show this help message")
	flag.StringVar(&template_path, "t", default_template_path, "Specify template path")
	flag.StringVar(&ddl_input_path, "ddl", "", "Specify the path of *.ddl.xml files")
	flag.StringVar(&gql_input_path, "gql", "", "Specify the path of *.gql.xml files")
	flag.StringVar(&output_path, "o", "./", "Specify output path")
	flag.Usage = usage
}

func ensure_dir(dir string) error {

	err := os.Mkdir(dir, os.ModeDir)

	if err == nil || os.IsExist(err) {
		return nil
	} else {
		return err
	}
}

func main() {
	flag.Parse()
	if help {
		flag.Usage()
		return
	}
	if !generator.DirExists(template_path) {
		log.Fatalf("Template path [%s] does not exists.", template_path)
		return
	}
	if !generator.DirExists(ddl_input_path) {
		log.Fatalf("*.ddl.xml path [%s] does not exists.", ddl_input_path)
		return
	}
	if !generator.DirExists(output_path) {
		log.Fatalf("Output path [%s] does not exists.", output_path)
		return
	}
	var p generator.Project
	p.Init()
	if err := p.LoadDDL(ddl_input_path); err != nil {
		return
	}
	if len(gql_input_path) != 0 {
		if generator.DirExists(gql_input_path) {
			if err := p.LoadGQL(gql_input_path); err != nil {
				log.Fatalf("load gql fail for [%s]", err)
			}
		} else {
			log.Fatalf("gql input path [%s] does not exists.", gql_input_path)
		}
	}
	p.Stat()
	if err := ensure_dir(path.Join(output_path, "gtable")); err != nil {
		log.Fatalf("Mkdir [%s] failed for %s", "gtable", err.Error())
		return
	}
	os.Chmod(path.Join(output_path, "gtable"), os.ModePerm)
	if err := ensure_dir(path.Join(output_path, "include")); err != nil {
		log.Fatalf("Mkdir [%s] failed for %s", "include", err.Error())
		return
	}
	os.Chmod(path.Join(output_path, "include"), os.ModePerm)
	if err := ensure_dir(path.Join(output_path, "src")); err != nil {
		log.Fatalf("Mkdir [%s] failed for %s", "src", err.Error())
		return
	}
	os.Chmod(path.Join(output_path, "src"), os.ModePerm)

	if err := p.Generate(template_path, output_path); err != nil {
		return
	}

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	cp := func(gtable_file string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		log.Printf(gtable_file)
		if !f.IsDir() {
			generator.Cp(gtable_file, path.Join(output_path, "gtable", filepath.Base(gtable_file)))
		}
		return nil
	}
	gtable_include := path.Join(filepath.Dir(ex), 
	    "../src/github.com/galois-advertising/gtable-generator/gtable-generator-include")
	filepath.Walk(gtable_include, cp)
	gtable_src := path.Join(filepath.Dir(ex), 
	    "../src/github.com/galois-advertising/gtable-generator/gtable-generator-src")
	filepath.Walk(gtable_src, cp)
	log.Println("Generate succeed.")
}
