//solopointer1202@gmail.com
package generator

import (
	"encoding/xml"
	"fmt"
	"log"
	"testing"
)

func TestSelect(t *testing.T) {
	blob := `
	<query>
    <name>find_xdv</name>
    <select>
      <column>user_table.user_id</column>
      <column>user_table.user_stat</column>
      <column>user_table.region</column>
      <column>plan_table.plan_id</column>
      <column>plan_table.region</column>
      <column>unit_table.unit_id</column>
      <column>xdv_table.xdv_id</column>
	</select>
	<from>
      <table join_type="FIRST" scan_limit="" result_limit="" each_scan_limit="" each_result_limit="">
        <name>xdv_table</name>
      </table>
      <table join_type="left join" scan_limit="" each_scan_limit="">
        <name>user_table</name>
        <on_column>xdv_table.user_id</on_column>
        <on_column>user_table.user_id</on_column>
      </table>
      <table join_type="left join" scan_limit="" each_scan_limit="">
        <name>plan_table</name>
        <on_column>xdv_table.plan_id</on_column>
        <on_column>plan_table.plan_id</on_column>
      </table>
      <table join_type="left join" scan_limit="" each_scan_limit="">
        <name>unit_table</name>
        <on_column>xdv_table.unit_id</on_column>
        <on_column>unit_table.unit_id</on_column>
      </table>
	</from>
	<where>
      <field_conditioner>
        <id>0</id>
        <type>=</type>
        <field>user_table.user_stat</field>
        <field>1</field>
      </field_conditioner>
      <field_conditioner>
        <id>1</id>
        <type>in</type>
        <field>user_table.region</field>
        <field>$TargetRegion</field>
      </field_conditioner>
      <field_conditioner>
        <id>3</id>
        <type>in</type>
        <field>plan_table.region</field>
        <field>$TargetRegion</field>
      </field_conditioner>
      <logic_conditioner>
        <id>2</id>
        <type>and</type>
        <sub_conditioner>0</sub_conditioner>
        <sub_conditioner>1</sub_conditioner>
      </logic_conditioner>
      <logic_conditioner>
        <id>4</id>
        <type>and</type>
        <sub_conditioner>2</sub_conditioner>
        <sub_conditioner>3</sub_conditioner>
      </logic_conditioner>
      <logic_conditioner>
        <id>5</id>
        <type>not</type>
        <sub_conditioner>4</sub_conditioner>
      </logic_conditioner>
    </where>
	</query>
	`
	var d Query
	if err := xml.Unmarshal([]byte(blob), &d); err != nil {
		log.Fatalf("Error:%s", err.Error())
	} else {
		log.Print(d)
	}
	fmt.Println(d.Name)
	fmt.Println(d.Columns)
	fmt.Println(d.Tables)
	fmt.Println(d.FieldConditions)
	fmt.Println(d.LogicConditions)
}
