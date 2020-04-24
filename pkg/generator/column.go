// solopointer1202@gmail.com
package generator

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
)

type Column struct {
	XMLName       xml.Name    `xml:"column_node"`
	Column_name   string      `xml:"name"`
	Column_kind   KindAttr    `xml:"kind"`
	IsDerivative string      `xml:"type,attr"`
	Constrains    []Constrain `xml:"constrains>constrain"`
	IsPrimarykey  bool
}

type KindAttr struct {
	Type   string `xml:"type,attr"`
	Length string `xml:"length,attr"`
	Kind   string `xml:",chardata"`
}

type Constrain struct {
	Prop string `xml:"prop,attr"`
	Name string `xml:",chardata"`
}

func (c *Column) Parse_from() (string, error) {
	if c.IsDerivative != "derivative" {
		return "", errors.New(fmt.Sprintf("%s is not a derivative column", c.Column_name))
	}
	for _, v := range c.Constrains {
		if v.Name == "from" {
			return v.Prop, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Could not find [from] of %s", c.Column_name))
}

func (c *Column) UpperName() string {
	return strings.ToUpper(c.Column_name)
}

func (c *Column) IsBasic() bool {
	return c.Column_kind.Kind == "basic" && c.Column_kind.Type != "binary"
}

func (c *Column) IsString() bool {
	return c.Column_kind.Kind == "array" && c.Column_kind.Type == "char"
}

func (c *Column) IsArray() bool {
	return c.Column_kind.Kind == "array" && c.Column_kind.Type != "char"
}

func (c *Column) Is_binary() bool {
	return c.Column_kind.Kind == "basic" && c.Column_kind.Type == "binary"
}

func (c *Column) Length() string {
	if c.IsArray() || c.IsString() {
		return c.Column_kind.Length
	} else {
		return "0"
	}
}
