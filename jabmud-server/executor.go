package main

import (
	"log"
)

type RunCommand func() // not sure args yet

var knownCommands = make(map[string]RunCommand)

func init() {
	knownCommands = map[string]RunCommand{
		"l":    Look,
		"look": Look,
	}
}

func Run(command string) {
	runner := knownCommands[command]
	if runner != nil {
		runner()
	} else {
		log.Printf("%s is not a known command", command)
	}
}

func Look() {
	log.Print("I looked")
}
