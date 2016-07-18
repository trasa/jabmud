package commands

import (
	"github.com/trasa/jabmud/world"
	"testing"
)

var center, north, south, west, east, up, down *world.Room
var player *world.Player

func TestMain(m *testing.M) {
	zone := &world.Zone{}
	center = world.NewRoom(zone, "center", "", "")
	north = world.NewRoom(zone, "north", "", "")
	south = world.NewRoom(zone, "south", "", "")
	west = world.NewRoom(zone, "west", "", "")
	east = world.NewRoom(zone, "east", "", "")
	up = world.NewRoom(zone, "up", "", "")
	down = world.NewRoom(zone, "down", "", "")

	center.North = north
	center.South = south
	center.East = east
	center.West = west
	center.Up = up
	center.Down = down

	player = &world.Player{
		Id: "foo",
	}
	center.AddPlayer(player)
}

func TestMoveNorthSuccess(t *testing.T) {

	result := MoveDirection(player, []string{"n"}).(MoveResult)
	if !result.Success {
		t.Error("should succeed")
	}

	if inRoom(center, player.Id) {
		t.Error("player found in center but should have moved")
	}
	if !inRoom(north, player.Id) {
		t.Error("shouldn't exist")
	}
	if inRoom(south, player.Id) {
		t.Error("shouldn't exist")
	}
	if inRoom(east, player.Id) {
		t.Error("shouldn't exist")
	}
	if inRoom(west, player.Id) {
		t.Error("shouldn't exist")
	}
	if inRoom(up, player.Id) {
		t.Error("shouldn't exist")
	}
	if inRoom(down, player.Id) {
		t.Error("shouldn't exist")
	}
}

func TestMoveSouthSuccess(t *testing.T) {

	result := MoveDirection(player, []string{"south"}).(MoveResult)
	if !result.Success {
		t.Error("should succeed")
	}
	if inRoom(center, player.Id) {
		t.Error("player found in center but should have moved")
	}
	if inRoom(north, player.Id) {
		t.Error("shouldn't exist")
	}
	if !inRoom(south, player.Id) {
		t.Error("shouldn't exist")
	}
	if inRoom(east, player.Id) {
		t.Error("shouldn't exist")
	}
	if inRoom(west, player.Id) {
		t.Error("shouldn't exist")
	}
	if inRoom(up, player.Id) {
		t.Error("shouldn't exist")
	}
	if inRoom(down, player.Id) {
		t.Error("shouldn't exist")
	}
}

func TestMoveEastSuccess(t *testing.T) {

	result := MoveDirection(player, []string{"e"}).(MoveResult)
	if !result.Success{
		t.Error("should succeed")

	}
	if inRoom(center, player.Id) {
		t.Error("player found in center but should have moved")
	}
	if inRoom(north, player.Id) {
		t.Error("shouldn't exist")
	}
	if inRoom(south, player.Id) {
		t.Error("shouldn't exist")
	}
	if !inRoom(east, player.Id) {
		t.Error("shouldn't exist")
	}
	if inRoom(west, player.Id) {
		t.Error("shouldn't exist")
	}
	if inRoom(up, player.Id) {
		t.Error("shouldn't exist")
	}
	if inRoom(down, player.Id) {
		t.Error("shouldn't exist")
	}
}

func TestMoveWestSuccess(t *testing.T) {

	result := MoveDirection(player, []string{"w"}).(MoveResult)
	if !result.Success {
		t.Error("should succeed")
	}
	if inRoom(center, player.Id) {
		t.Error("player found in center but should have moved")
	}
	if inRoom(north, player.Id) {
		t.Error("shouldn't exist")
	}
	if inRoom(south, player.Id) {
		t.Error("shouldn't exist")
	}
	if inRoom(east, player.Id) {
		t.Error("shouldn't exist")
	}
	if !inRoom(west, player.Id) {
		t.Error("shouldn't exist")
	}
	if inRoom(up, player.Id) {
		t.Error("shouldn't exist")
	}
	if inRoom(down, player.Id) {
		t.Error("shouldn't exist")
	}
}

func TestMoveUpSuccess(t *testing.T) {

	result := MoveDirection(player, []string{"w"}).(MoveResult)
	if !result.Success {
		t.Error("should succeed")
	}
	if inRoom(center, player.Id) {
		t.Error("player found in center but should have moved")
	}
	if inRoom(north, player.Id) {
		t.Error("shouldn't exist")
	}
	if inRoom(south, player.Id) {
		t.Error("shouldn't exist")
	}
	if inRoom(east, player.Id) {
		t.Error("shouldn't exist")
	}
	if inRoom(west, player.Id) {
		t.Error("shouldn't exist")
	}
	if !inRoom(up, player.Id) {
		t.Error("shouldn't exist")
	}
	if inRoom(down, player.Id) {
		t.Error("shouldn't exist")
	}
}

func TestMoveDownSuccess(t *testing.T) {

	result := MoveDirection(player, []string{"w"}).(MoveResult)
	if !result.Success {
		t.Error("should succeed")
	}
	if inRoom(center, player.Id) {
		t.Error("player found in center but should have moved")
	}
	if inRoom(north, player.Id) {
		t.Error("shouldn't exist")
	}
	if inRoom(south, player.Id) {
		t.Error("shouldn't exist")
	}
	if inRoom(east, player.Id) {
		t.Error("shouldn't exist")
	}
	if inRoom(west, player.Id) {
		t.Error("shouldn't exist")
	}
	if inRoom(up, player.Id) {
		t.Error("shouldn't exist")
	}
	if !inRoom(down, player.Id) {
		t.Error("shouldn't exist")
	}
}

func inRoom(room *world.Room, id string) (exists bool) {
	_, exists = room.Players[id]
	return
}
