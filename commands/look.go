package commands

import (
	"fmt"
	"github.com/trasa/jabmud/world"
	"log"
)

type LookResult struct {
	Value     string
	Title     string `xml:"Title,omitempty"`
	PlayerIds []string
}

func (lr LookResult) String() string {
	return fmt.Sprintf("(LookResult Title='%s', Value='%s', PlayerIds='%v')", lr.Title, lr.Value, lr.PlayerIds)
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
		playerIds := make([]string, len(playerRoom.Players))
		for _, p := range playerRoom.Players {
			log.Printf("found player %s", p.Id)
			playerIds = append(playerIds, p.Id)
		}
		return LookResult{
			Title:     playerRoom.Name,
			Value:     playerRoom.Description,
			PlayerIds: playerIds,
		}
	}
}
