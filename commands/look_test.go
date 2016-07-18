package commands

import (
	"github.com/trasa/jabmud/world"
	"log"
	"testing"
)

func TestLookPlayerNotInRoom(t *testing.T) {
	p := world.Player{}

	result := Look(&p, nil).(LookResult)
	log.Printf("%s", result)
	if len(result.PlayerIds) != 0 {
		t.Error("room playerids result should be empty")
	}
}

func TestLookPlayers(t *testing.T) {
	room := world.NewRoom(nil, "id", "name", "desc")
	p := world.Player{
		Id: "foo",
	}
	room.AddPlayer(&p)

	result := Look(&p, nil).(LookResult)
	log.Printf("result %v", result)

	if len(result.PlayerIds) != 1 {
		t.Error("too many players")
	}
	if result.PlayerIds[0] != "foo" {
		t.Error("didnt find player")
	}
}
