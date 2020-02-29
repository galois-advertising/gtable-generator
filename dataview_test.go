package main

import (
	"encoding/xml"
	"log"
	"os"
	"testing"
)

func TestDataview(t *testing.T) {
	blob := `
	<dataview>
    <name>UserView</name>
    <channel>UserSource::UserTable</channel>
    <properties>
      <udf>UserViewUDF</udf>
    </properties>
    <columns_node>
      <column_node type="original">
        <name>user_id</name>
        <type>uint64</type>
        <constrains/>
      </column_node>
      <column_node type="original">
        <name>adx_id</name>
        <type>uint64</type>
        <constrains/>
      </column_node>
      <column_node type="original">
        <name>wise_vectors</name>
        <type type="char" length="Constant::MAX_BUF_LEN">array</type>
        <constrains>
          <constrain>opt</constrain>
        </constrains>
      </column_node>
      <column_node type="original">
        <name>pc_vectors</name>
        <type type="char" length="Constant::MAX_BUF_LEN">array</type>
        <constrains>
          <constrain>opt</constrain>
        </constrains>
      </column_node>
      <column_node type="derivative">
        <name>usr_id_2</name>
        <type>uint32</type>
        <constrains>
          <constrain>del</constrain>
          <constrain prop="user_id,adx_id">from</constrain>
        </constrains>
      </column_node>
      <column_node type="derivative">
        <name>black_wd</name>
        <type type="uint64" length="Constant::MAX_BLACK_WORD_NUM">array</type>
        <constrains>
          <constrain prop="adx_id">from</constrain>
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
		if res, err := xml.MarshalIndent(d, "", "    "); err != nil {
			log.Fatalf("Error:%s", err.Error())
		} else {
			os.Stdout.Write(res)
		}
	}
}
