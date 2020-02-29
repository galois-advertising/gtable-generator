//solopointer1202@gmail.com
package main

import (
	"log"
	"testing"
)

func TestProject(t *testing.T) {
	var p Project
	p.Init()
	p.LoadDDL("./")
	p.Stat()
	if err := p.Generate("../../templates", "../../output"); err == nil {
		log.Println("succeed")
	}
}
