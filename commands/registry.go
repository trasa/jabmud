package commands

import "log"

type RunCommand func(args []string) interface{}

var knownCommands = make(map[string]RunCommand)

func init() {
	knownCommands = map[string]RunCommand{
		"l":     Look,
		"login": Login,
		"look":  Look,
	}
}

// Run the command identified in the knownCommands registry.
// If the command (or an alias) isn't found, raises an error.
func Run(command string, args []string) interface{} {
	runner := knownCommands[command]
	if runner != nil {
		return runner(args)
	} else {
		log.Printf("%s is not a known command", command)
		return nil
	}
}

type LookResult struct {
	Value string
}

type LoginResult struct {
	Success bool
}

// Look around you.
func Look(args []string) interface{} {
	log.Printf("I looked: %s", args)
	return LookResult{"You don't see anything."}
}

func Login(args []string) interface{} {
	log.Print("Login %s", args)
	return LoginResult{true}
}
