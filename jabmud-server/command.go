package main

import (
	"encoding/xml"
	"fmt"
)

type Command struct {
	Name    string   `xml:"cmdName,attr"`
	ArgList []string `xml:"arg"`
}

func (c Command) String() string {
	s := fmt.Sprintf("<command cmdName='%s'>", c.Name)
	for _, ca := range c.ArgList {
		s += fmt.Sprintf("<arg>%s</arg>", ca)
	}
	s += "</command>"
	return s
}

func ParseCommand(raw string) Command {
	var cmd Command
	xml.Unmarshal([]byte(raw), &cmd)
	return cmd
}
