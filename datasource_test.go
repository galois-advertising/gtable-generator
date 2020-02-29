package main

import (
	"encoding/xml"
	"log"
	"testing"
)

func TestDatasource(t *testing.T) {
	blob := `
  <datasource>
    <name>gbus</name>
    <properties>
      <type>databus</type>
      <protourl>ssh://git@github.com/galois-advertising/common/master/dbschema/freyja/databus_event.proto</protourl>
    </properties>
    <notations/>
  </datasource>
	`
	var d Datasource
	if err := xml.Unmarshal([]byte(blob), &d); err != nil {
		log.Fatalf("Error:%s", err.Error())
	} else {
		log.Print(d)
	}
}
