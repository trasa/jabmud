package jab

import (
	"encoding/xml"
	"fmt"
)

type IqCommand struct {
	Name    string   `xml:"name,attr"`
	ArgList []string `xml:"arg"`
}

func (c IqCommand) String() string {
	s := fmt.Sprintf("<command name='%s'>", c.Name)
	for _, ca := range c.ArgList {
		s += fmt.Sprintf("<arg>%s</arg>", ca)
	}
	s += "</command>"
	return s
}

func DeserializeIqCommand(rawxml string) IqCommand {
	var cmd IqCommand
	xml.Unmarshal([]byte(rawxml), &cmd)
	return cmd
}
