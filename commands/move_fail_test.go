package commands

import (
	"github.com/trasa/jabmud/world"
	"testing"
)

var center *world.Room
var player *world.Player

func TestMain(m *testing.M) {
	zone := &world.Zone{}
	center = world.NewRoom(zone, "center", "", "")

	player = &world.Player{
		Id: "foo",
	}
	center.AddPlayer(player)
}

func TestMoveBadDirection(t *testing.T) {
	result := MoveDirection(player, []string{"x"}).(MoveResult)
	if result.Success {
		t.Error("should fail")
	}
}
