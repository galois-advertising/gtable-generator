//solopointer1202@gmail.com
package generator

import (
	"encoding/xml"
	"fmt"
	"log"
	"testing"
)

func TestDataupdator(t *testing.T) {
	blob := `
	<dataupdator>
    <name>plan_view|plan_table</name>
    <from>plan_view</from>
    <to>plan_table</to>
    <properties>
      <udf>plan_view_to_plan_table</udf>
      <type>NOT DEFAULT</type>
    </properties>
    <notations/>
  </dataupdator>
	`
	var d Dataupdator
	if err := xml.Unmarshal([]byte(blob), &d); err != nil {
		log.Fatalf("Error:%s", err.Error())
	} else {
		log.Print(d)
	}
	d.Setup()
	fmt.Print(d)
}
