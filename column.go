// solopointer1202@gmail.com
package main

import (
	"encoding/xml"
	"strings"
)

type Column struct {
	XMLName         xml.Name    `xml:"column_node"`
	Column_name     string      `xml:"name"`
	Column_kind     KindAttr    `xml:"kind"`
	Colume_from     string      `xml:"type,attr"`
	Constrains_list []Constrain `xml:"constrains>constrain"`
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

func (c *Column) UpperName() string {
	return strings.ToUpper(c.Column_name)
}

func (c *Column) Is_basic() bool {
	return c.Column_kind.Kind == "basic" && c.Column_kind.Type != "binary"
}

func (c *Column) Is_string() bool {
	return c.Column_kind.Kind == "array" && c.Column_kind.Type == "char"
}

func (c *Column) Is_array() bool {
	return c.Column_kind.Kind == "array" && c.Column_kind.Type != "char"
}

func (c *Column) Is_binary() bool {
	return c.Column_kind.Kind == "basic" && c.Column_kind.Type == "binary"
}

func (c *Column) Length() string {
	if c.Is_array() || c.Is_string() {
		return c.Column_kind.Length
	} else {
		return "0"
	}
}
