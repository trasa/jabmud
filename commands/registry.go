package commands

import (
	"encoding/xml"
	"log"
	"github.com/trasa/jabmud/world"
)

type RunCommand func(player world.Player, args []string) interface{}

var knownCommands = make(map[string]RunCommand)

func init() {
	knownCommands = map[string]RunCommand{
		"l":    Look,
		"look": Look,
		"who":  Who,
	}
}

// Run the command identified in the knownCommands registry.
// If the command (or an alias) isn't found, raises an error.
func Run(player world.Player, command string, args []string) interface{} {
	runner := knownCommands[command]
	if runner != nil {
		return runner(player, args)
	} else {
		log.Printf("%s is not a known command", command)
		return nil
	}
}

// Serialize obj into it's xml representation as a string.
// If obj is nil, return empty-string.
func Serialize(obj interface{}) string {
	if obj == nil {
		return ""
	}
	bytes, _ := xml.Marshal(obj)
	return string(bytes)
}

type LookResult struct {
	Value string
}

// Look around you.
func Look(player world.Player, args []string) interface{} {
	log.Printf("%s looked: %s", player , args)
	return LookResult{"You don't see anything."}
}

