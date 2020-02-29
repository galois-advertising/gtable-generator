package main

import (
	"encoding/xml"
	"log"
	"testing"
)

func TestIndextable(t *testing.T) {
	blob := `
  <indextable>
    <name>user_index</name>
    <on_table>user_table</on_table>
    <on_column>region</on_column>
    <properties>
      <type>HashTable</type>
    </properties>
    <columns_node>
      <column_node type="original">
        <name>user_id</name>
        <kind type="uint32_t">basic</kind>
        <constrains/>
      </column_node>
      <column_node type="original">
        <name>user_stat</name>
        <kind type="uint32_t">basic</kind>
        <constrains/>
      </column_node>
      <column_node type="original">
        <name>region</name>
        <kind type="uint64_t">basic</kind>
        <constrains/>
      </column_node>
    </columns_node>
    <notations/>
  </indextable>
	`
	var d Indextable
	if err := xml.Unmarshal([]byte(blob), &d); err != nil {
		log.Fatalf("Error:%s", err.Error())
	} else {
		log.Print(d)
	}
}
