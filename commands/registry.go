package commands

import "log"

type RunCommand func(args []string)

var knownCommands = make(map[string]RunCommand)

func init() {
	knownCommands = map[string]RunCommand{
		"l":    Look,
		"look": Look,
	}
}

// Run the command identified in the knownCommands registry.
// If the command (or an alias) isn't found, raises an error.
func Run(command string, args []string) {
	runner := knownCommands[command]
	if runner != nil {
		runner(args)
	} else {
		log.Printf("%s is not a known command", command)
	}
}

// Look around you.
// TODO move this to some other file full of command implementations
func Look(args []string) {
	log.Printf("I looked: %s", args)
}
