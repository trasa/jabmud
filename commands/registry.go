package commands

import (
	"github.com/trasa/jabmud/world"
	"log"
)

type RunCommand func(player *world.Player, args []string) interface{}

var knownCommands = make(map[string]RunCommand)

func init() {
	knownCommands = map[string]RunCommand{
		"l":    Look,
		"look": Look,
		"m":    MoveDirection,
		"move": MoveDirection,
		"who":  Who,
	}
}

// Run the command identified in the knownCommands registry.
// If the command (or an alias) isn't found, raises an error.
func Run(player *world.Player, command string, args []string) interface{} {
	runner := knownCommands[command]
	if runner != nil {
		return runner(player, args)
	} else {
		log.Printf("%s is not a known command", command)
		return nil
	}
}
