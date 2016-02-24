package commands

import (
	"github.com/trasa/jabmud/world"
	"log"
)

type WhoResult struct {
	// TODO this might have to become more sophisticated one day.
	// just not today.
	Players []*world.Player `xml:"Player"`
}

// Who else is online?
func Who(player *world.Player, args []string) interface{} {
	log.Printf("%s wants to know who is online", player)
	players := world.GetAllPlayers()
	return WhoResult{players}
}
