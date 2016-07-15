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
	room := world.Room{}
	p := world.Player{
		Id: "foo",
	}
	room.AddPlayer(&p)

	result := Look(&p, nil).(LookResult)
	log.Printf("result %v", result)

	foundIt := false
	for _, id := range result.PlayerIds {
		if id == "foo" {
			foundIt = true
			break
		}
	}
	if !foundIt {
		t.Error("should have found player id in look result")
	}

}
