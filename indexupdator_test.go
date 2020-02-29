//solopointer1202@gmail.com
package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"testing"
)

func TestIndexupdator(t *testing.T) {
	blob := `
	<indexupdator>
    <name>user_table|user_index</name>
    <from>user_table</from>
    <to>user_index</to>
    <properties>
      <type>NOT DEFAULT</type>
    </properties>
    <notations/>
  </indexupdator>
	`
	var d Indexupdator
	if err := xml.Unmarshal([]byte(blob), &d); err != nil {
		log.Fatalf("Error:%s", err.Error())
	} else {
		log.Print(d)
	}
	d.Setup()
	fmt.Print(d)
}
