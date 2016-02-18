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

// Tells the server that you are here and we should note your presence.
// I'd rather use something built-in, like, say, 'presence' but haven't
// been able to get that to function between client, ejabberd, and this
// component. So leaving that as a TODO.
func Login(args []string) interface{} {
	log.Print("Login %s", args)
	//	playerName := args[0]
	// if the player isn't in the game world, put them there in the start room
	// some other boring setup of their character...
	return LoginResult{true}
}
