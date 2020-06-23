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
      <name>SeekIndex</name>
      <select>
        <column>AdxTable.Adx_id</column>
        <column>AdxTable.usr_id</column>
      </select>
      <from>
        <table join_type="FIRST" scan_limit="$ScanLimit" result_limit="$Limit" each_scan_limit="" each_result_limit="">
          <name>Index</name>
        </table>
        <table join_type="join" scan_limit="" each_scan_limit="">
          <name>AdxTable</name>
          <on_column>Index.Adx_id</on_column>
          <on_column>AdxTable.Adx_id</on_column>
        </table>
        <table join_type="left join" scan_limit="" each_scan_limit="">
          <name>User</name>
          <on_column>AdxTable.plan_id</on_column>
          <on_column>$Version</on_column>
          <on_column>User.plan_id</on_column>
          <on_column>User.version</on_column>
        </table>
        <table join_type="join" scan_limit="" each_scan_limit="">
          <name>PlanTable</name>
          <on_column>AdxTable.plan_id</on_column>
          <on_column>PlanTable.plan_id</on_column>
        </table>
      </from>
      <where>
        <field_conditioner>
          <id>0</id>
          <type>=</type>
          <field>Index.exact_sign</field>
          <field>$KeySign</field>
        </field_conditioner>
        <field_conditioner>
          <id>1</id>
          <type>=</type>
          <field>AdxTable.new_match_type</field>
          <field>$match_type1</field>
        </field_conditioner>
        <field_conditioner>
          <id>2</id>
          <type>=</type>
          <field>AdxTable.new_match_type</field>
          <field>$match_type2</field>
        </field_conditioner>
        <field_conditioner>
          <id>4</id>
          <type>=</type>
          <field>AdxTable.new_match_type</field>
          <field>$match_type3</field>
        </field_conditioner>
        <logic_conditioner>
          <id>3</id>
          <type>or</type>
          <sub_conditioner>1</sub_conditioner>
          <sub_conditioner>2</sub_conditioner>
        </logic_conditioner>
        <logic_conditioner>
          <id>5</id>
          <type>or</type>
          <sub_conditioner>3</sub_conditioner>
          <sub_conditioner>4</sub_conditioner>
        </logic_conditioner>
        <logic_conditioner>
          <id>6</id>
          <type>and</type>
          <sub_conditioner>0</sub_conditioner>
          <sub_conditioner>5</sub_conditioner>
        </logic_conditioner>
      </where>
    </query>
    <query>
      <name>FindWithAdxid</name>
      <select>
        <column>AdxTable.adx_id</column>
        <column>UserTable.usr_id</column>
        <column>PlanTable.plan_id</column>
        <column>UnitTable.unit_id</column>
      </select>
      <from>
        <table join_type="FIRST" scan_limit="" result_limit="" each_scan_limit="" each_result_limit="">
          <name>AdxTable</name>
        </table>
        <table join_type="join" scan_limit="" each_scan_limit="">
          <name>PlanTable</name>
          <on_column>AdxTable.plan_id</on_column>
          <on_column>PlanTable.plan_id</on_column>
        </table>
        <table join_type="join" scan_limit="" each_scan_limit="">
          <name>UserTable</name>
          <on_column>AdxTable.usr_id</on_column>
          <on_column>UserTable.usr_id</on_column>
        </table>
        <table join_type="join" scan_limit="" each_scan_limit="">
          <name>UnitTable</name>
          <on_column>AdxTable.unit_id</on_column>
          <on_column>UnitTable.unit_id</on_column>
        </table>
      </from>
      <where>
        <field_conditioner>
          <id>0</id>
          <type>in</type>
          <field>AdxTable.adx_id</field>
          <field>$adx_id_list</field>
        </field_conditioner>
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
