package generator

import (
	"encoding/xml"
	"log"
	"testing"
)

func TestDatatable(t *testing.T) {
	blob := `
  <datatable>
    <name>UserTable6</name>
    <properties>
      <type>HashTable</type>
      <hash_ratio>0.5</hash_ratio>
    </properties>
    <columns_node>
      <column_node type="original">
        <name>usr_id_1</name>
        <kind type="uint8">basic</kind>
        <constrains>
          <constrain prop="1,2,3,4,5">range</constrain>
          <constrain>opt</constrain>
        </constrains>
      </column_node>
      <column_node type="original">
        <name>usr_id_2</name>
        <kind type="uint16">basic</kind>
        <constrains>
          <constrain>opt</constrain>
        </constrains>
      </column_node>
      <column_node type="original">
        <name>usr_id_3</name>
        <kind type="uint32">basic</kind>
        <constrains>
          <constrain>del</constrain>
        </constrains>
      </column_node>
      <column_node type="original">
        <name>usr_id_4</name>
        <kind type="uint64" length="10u">array</kind>
        <constrains/>
      </column_node>
      <column_node type="original">
        <name>usr_id_5</name>
        <kind type="bool" length="MAX_LEN">array</kind>
        <constrains/>
      </column_node>
      <column_node type="original">
        <name>usr_id_6</name>
        <kind type="char" length="12u">array</kind>
        <constrains/>
      </column_node>
      <column_node type="original">
        <name>usr_id_7</name>
        <kind type="int">basic</kind>
        <constrains>
          <constrain prop="11">default</constrain>
        </constrains>
      </column_node>
      <column_node type="original">
        <name>usr_id_8</name>
        <kind type="long">basic</kind>
        <constrains>
          <constrain prop="uid8">as</constrain>
        </constrains>
      </column_node>
      <column_node type="original">
        <name>usr_id_9</name>
        <kind type="float">basic</kind>
        <constrains>
          <constrain>custom</constrain>
        </constrains>
      </column_node>
      <column_node type="original">
        <name>usr_id_10</name>
        <kind type="binary">basic</kind>
        <constrains/>
      </column_node>
    </columns_node>
    <primary_key type="uint64key">usr_id_10,usr_id_9</primary_key>
    <notations>
      <notation>Galois Advertising Framework</notation>
      <notation>solopointer</notation>
    </notations>
  </datatable>
	`
	var d Datatable
	if err := xml.Unmarshal([]byte(blob), &d); err != nil {
		log.Fatalf("Error:%s", err.Error())
	} else {
		log.Print(d)
	}
}
