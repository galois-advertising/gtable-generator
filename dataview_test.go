package main

import (
	"encoding/xml"
	"log"
	"os"
	"testing"
	"text/template"
)

func TestDataview(t *testing.T) {
	blob := `
  <dataview>
  <name>user_view</name>
  <channel>gbus::user_event</channel>
  <property name="udf">user_view_udf</property>
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
    <column_node type="original">
      <name>user_name</name>
      <kind type="char" length="1024u">array</kind>
      <constrains/>
    </column_node>
    <column_node type="derivative">
      <name>user_name_sign</name>
      <kind type="uint64_t">basic</kind>
      <constrains>
        <constrain prop="user_name">from</constrain>
      </constrains>
    </column_node>
  </columns_node>
  <notations/>
</dataview> 
	`
	var d Dataview
	if err := xml.Unmarshal([]byte(blob), &d); err != nil {
		log.Fatalf("Error:%s", err.Error())
	} else {
		const tmpl = `
        {{.Name}}
        {{.GetUDF}}
    `
		t := template.Must(template.New("html").Parse(tmpl))
		err = t.Execute(os.Stdout, &d)
		if err != nil {
			panic(err)
		}
		//if res, err := xml.MarshalIndent(d, "", "    "); err != nil {
		//	log.Fatalf("Error:%s", err.Error())
		//} else {
		//	os.Stdout.Write(res)
		//}
	}
}
