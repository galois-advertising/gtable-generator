package main

import (
	"ddldefines"
	"flag"
	"fmt"
	"gqldefines"
	"os"
	"strings"
)

func help() {
	fmt.Fprintf(os.Stderr, "Usage:  %s [-ddl *.ddl] [-gql *.gql] [-output path]\nOptions:\n",
		os.Args[0])
	flag.PrintDefaults()
}

func main() {
	var ddl = flag.String("ddl", "", "Input a *.ddl file.")
	var gql = flag.String("gql", "", "Input a *.gql file.")
	var output = flag.String("output", "", "The output path.")
	flag.Usage = help
	flag.Parse()
	fmt.Print(*ddl, *gql, *output)
	var s string = "hello world!"
	fmt.Println(strings.Replace(s, "hello", "Hello", 1))
	if *ddl != "" {
		dd := ddldefines.DdlDefines{dataviews: make(map[string]string)}
		fmt.Println(dd.load(*ddl))
	}
	if *gql != "" {
		gd := gqldefines.GqlDefines{}
		fmt.Println(gd.load(*gql))
	}

}
