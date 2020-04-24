package generator

import (
	"encoding/xml"
	"log"
	"testing"
)

func TestField(t *testing.T) {
	blob := `
	<column_node type="derivative">
		<name>usr_id_2</name>
		<type type="char" length="Constant::MAX_BUF_LEN">array</type>
        <constrains>
          <constrain>del</constrain>
          <constrain prop="user_id,adx_id">from</constrain>
        </constrains>
	</column_node>
	`
	var f Column
	if err := xml.Unmarshal([]byte(blob), &f); err != nil {
		log.Fatalf("Error:%s", err.Error())
	} else {
		log.Print(f)
		log.Println(f.Constrains[1].Prop)
		log.Println(len(f.Constrains[0].Prop))
	}
}
