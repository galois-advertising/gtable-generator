package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:  %s [-t templates_dir] [-i ddl_path] [-o path]\nOptions:\n",
		os.Args[0])
	flag.PrintDefaults()
}

var (
	help          bool
	template_path string
	input_path    string
	output_path   string
)

func init() {
	var default_template_path string
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	default_template_path = path.Join(filepath.Dir(ex), "gtable-generator-templates")
	flag.BoolVar(&help, "h", false, "Show this help message")
	flag.StringVar(&template_path, "t", default_template_path, "Specify template path")
	flag.StringVar(&input_path, "i", "", "Specify the path of *.ddl.xml files")
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
	if !DirExists(template_path) {
		log.Fatalf("Template path [%s] does not exists.", template_path)
		return
	}
	if !DirExists(input_path) {
		log.Fatalf("*.ddl.xml path [%s] does not exists.", input_path)
		return
	}
	if !DirExists(output_path) {
		log.Fatalf("Output path [%s] does not exists.", output_path)
		return
	}
	var p Project
	p.Init()
	if err := p.LoadDDL(input_path); err != nil {
		return
	}
	p.Stat()
	log.Println(path.Join(output_path, "gtable"))
	if err := ensure_dir(path.Join(output_path, "gtable")); err != nil {
		log.Fatalf("Mkdir [%s] failed for %s", "gtable", err.Error())
		return
	}
	if err := ensure_dir(path.Join(output_path, "include")); err != nil {
		log.Fatalf("Mkdir [%s] failed for %s", "include", err.Error())
		return
	}
	if err := ensure_dir(path.Join(output_path, "src")); err != nil {
		log.Fatalf("Mkdir [%s] failed for %s", "src", err.Error())
		return
	}
	if err := p.Generate(template_path, output_path); err != nil {
		return
	}
	log.Println("Generate succeed.")
}
