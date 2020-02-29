package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func help() {
	fmt.Fprintf(os.Stderr, "Usage:  %s [-t templates_dir] [-i ddl_path] [-o path]\nOptions:\n",
		os.Args[0])
	flag.PrintDefaults()
}

func main() {
	var templates_path = flag.String("t", "", "Templates file path.")
	var ddl_path = flag.String("i", "", "Input ddl file path.")
	var output_path = flag.String("o", "", "The output path.")
	flag.Usage = help
	flag.Parse()
	fmt.Print(*templates_path, *ddl_path, *output_path)
	if !DirExists(*templates_path) {
		log.Fatalf("Template path [%s] does not exists.", *templates_path)
		return
	}
	if !DirExists(*ddl_path) {
		log.Fatalf("DDL path [%s] does not exists.", *ddl_path)
		return
	}
	if !DirExists(*output_path) {
		log.Fatalf("Output path [%s] does not exists.", *output_path)
		return
	}
	var p Project
	p.Init()
	p.LoadDDL(*ddl_path)
	p.Stat()
	p.Generate(*templates_path, *output_path)
}
