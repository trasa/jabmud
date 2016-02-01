package main

import (
	"log"
	"testing"
)

func TestParseCommand(t *testing.T) {
	s := "<command cmdName='blargh'></command>"
	cmd := ParseCommand(s)
	log.Printf("cmd: %s", cmd)

	if cmd.Name != "blargh" {
		t.Errorf("cmd not named blargh: %s", cmd.Name)
	}
	if cmd.ArgList != nil {
		t.Error("found an arglist where it wasn't expected")
	}
}

func TestParseCommandArgList(t *testing.T) {
	s := "<command cmdName='blargh'><arg>one</arg><arg>two</arg></command>"
	cmd := ParseCommand(s)
	log.Printf("cmd: %s", cmd)

	if cmd.Name != "blargh" {
		t.Errorf("cmd not named blargh: %s", cmd.Name)
	}

	if cmd.ArgList == nil || len(cmd.ArgList) == 0 {
		t.Error("missing arglist")
	}
	if cmd.ArgList[0] != "one" {
		t.Error("expected one")
	}
	if cmd.ArgList[1] != "two" {
		t.Error("expected two")
	}
}
