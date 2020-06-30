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
  <?xml version="1.0" encoding="UTF-8"?>
<gql>
  <parser_build_time>Jun 27 2020 13:41:52</parser_build_time>
  <handler>demo</handler>
  <namespace>galois::user</namespace>
  <parser_build_time>Jun 27 2020 13:41:52</parser_build_time>
  <query>
    <name>SeekIndex</name>
    <select>
      <column>AdxTable.Adx_id</column>
      <column>AdxTable.usr_id</column>
    </select>
    <from>
      <table join_type="base">
        <name>Index</name>
        <scan_limit>$scanlimit</scan_limit>
        <result_limit>$limit</result_limit>
      </table>
      <table join_type="join">
        <name>AdxTable</name>
        <scan_limit>$scanlimit2</scan_limit>
        <each_scan_limit>$eachscanlimit2</each_scan_limit>
        <left_on_columns>
          <field>Index.Adx_id</field>
        </left_on_columns>
        <right_on_columns>
          <field>AdxTable.Adx_id</field>
        </right_on_columns>
      </table>
      <table join_type="leftjoin">
        <name>User</name>
        <scan_limit>$scanlimit3</scan_limit>
        <left_on_columns>
          <field>AdxTable.plan_id</field>
          <field>$Version</field>
        </left_on_columns>
        <right_on_columns>
          <field>User.plan_id</field>
          <field>User.version</field>
        </right_on_columns>
      </table>
      <table join_type="join">
        <name>PlanTable</name>
        <scan_limit>$scanlimit4</scan_limit>
        <left_on_columns>
          <field>AdxTable.plan_id</field>
        </left_on_columns>
        <right_on_columns>
          <field>PlanTable.plan_id</field>
        </right_on_columns>
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
      <field_conditioner>
        <id>7</id>
        <type>!=</type>
        <field apply="@len">AdxTable.url_sign</field>
        <field>$zero</field>
      </field_conditioner>
      <field_conditioner>
        <id>9</id>
        <type>&gt;=</type>
        <field apply="@max">AdxTable.bid</field>
        <field>$one</field>
      </field_conditioner>
      <field_conditioner>
        <id>11</id>
        <type>=</type>
        <field apply="@not_null">PlanTable.plan_id</field>
        <field>$true</field>
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
      <logic_conditioner>
        <id>8</id>
        <type>and</type>
        <sub_conditioner>6</sub_conditioner>
        <sub_conditioner>7</sub_conditioner>
      </logic_conditioner>
      <logic_conditioner>
        <id>10</id>
        <type>and</type>
        <sub_conditioner>8</sub_conditioner>
        <sub_conditioner>9</sub_conditioner>
      </logic_conditioner>
      <logic_conditioner>
        <id>12</id>
        <type>and</type>
        <sub_conditioner>10</sub_conditioner>
        <sub_conditioner>11</sub_conditioner>
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
      <table join_type="base">
        <name>AdxTable</name>
      </table>
      <table join_type="join">
        <name>PlanTable</name>
        <left_on_columns>
          <field>AdxTable.plan_id</field>
        </left_on_columns>
        <right_on_columns>
          <field>PlanTable.plan_id</field>
        </right_on_columns>
      </table>
      <table join_type="join">
        <name>UserTable</name>
        <left_on_columns>
          <field>AdxTable.usr_id</field>
        </left_on_columns>
        <right_on_columns>
          <field>UserTable.usr_id</field>
        </right_on_columns>
      </table>
      <table join_type="join">
        <name>UnitTable</name>
        <left_on_columns>
          <field>AdxTable.unit_id</field>
        </left_on_columns>
        <right_on_columns>
          <field>UnitTable.unit_id</field>
        </right_on_columns>
      </table>
    </from>
    <where>
      <field_conditioner>
        <id>0</id>
        <type>in</type>
        <field>AdxTable.adx_id</field>
        <field>$adx_id_list</field>
      </field_conditioner>
      <field_conditioner>
        <id>1</id>
        <type>in</type>
        <field>AdxTable.plan_id</field>
        <field>$planid_list</field>
      </field_conditioner>
      <field_conditioner>
        <id>3</id>
        <type>in</type>
        <field>AdxTable.owmatch</field>
        <field>$owmatch_exclude_list</field>
      </field_conditioner>
      <field_conditioner>
        <id>4</id>
        <type>&gt;</type>
        <field>UserTable.usr_id</field>
        <field>$max_user_id</field>
      </field_conditioner>
      <logic_conditioner>
        <id>2</id>
        <type>and</type>
        <sub_conditioner>0</sub_conditioner>
        <sub_conditioner>1</sub_conditioner>
      </logic_conditioner>
      <logic_conditioner>
        <id>5</id>
        <type>or</type>
        <sub_conditioner>3</sub_conditioner>
        <sub_conditioner>4</sub_conditioner>
      </logic_conditioner>
      <logic_conditioner>
        <id>6</id>
        <type>not</type>
        <sub_conditioner>5</sub_conditioner>
      </logic_conditioner>
      <logic_conditioner>
        <id>7</id>
        <type>and</type>
        <sub_conditioner>2</sub_conditioner>
        <sub_conditioner>6</sub_conditioner>
      </logic_conditioner>
    </where>
  </query>
</gql>
	`
	var ds GqlDefines
	if err := xml.Unmarshal([]byte(blob), &ds); err != nil {
		log.Fatalf("Error:%s", err.Error())
	} else {
		ds.Setup()
	}
	fmt.Println(ds.Handler)
	fmt.Println(ds.Namespace)
	fmt.Println(ds.ParserBuildTime)
	for _, d := range ds.Queries {

		fmt.Println(d.Name)
		fmt.Println("")
		for _, col := range d.Columns {
			fmt.Println(col)
		}
		fmt.Println("")
		for _, tab := range d.Tables {
			fmt.Println(tab)
		}
		fmt.Println("")
		for _, field := range d.FieldConditions {
			fmt.Println(field)
		}
		fmt.Println("")
		for _, field := range d.LogicConditions {
			fmt.Println(field)
		}
	}
}
