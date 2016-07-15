package commands

import (
	"github.com/trasa/jabmud/world"
	"log"
)

type LookResult struct {
	Value string
	Title string `xml:"Title,omitempty"`
}

// Look around you.
func Look(player *world.Player, args []string) interface{} {
	log.Printf("%s looked: %s", player, args)

	playerRoom := player.Room
	if playerRoom == nil {
		return LookResult{
			Title: "Not In A Room",
			Value: "You see nothing but endless void.",
		}
	} else {
		return LookResult{
			Title: playerRoom.Name,
			Value: playerRoom.Description,
		}
	}
}
