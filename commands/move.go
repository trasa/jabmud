package commands

import (
	"github.com/trasa/jabmud/world"
	"log"
	"strings"
)

type MoveResult struct {
	Success bool
	Result  string
}

// Move from here to there
func MoveDirection(player *world.Player, args []string) interface{} {

	dir := args[0]
	dir = strings.TrimSpace(strings.ToLower(dir))
	var destinationRoom *world.Room

	if strings.HasPrefix(dir, "n") {
		// north
		destinationRoom = player.Room.North
	} else if strings.HasPrefix(dir, "s") {
		// south
		destinationRoom = player.Room.South
	} else if strings.HasPrefix(dir, "e") {
		// east
		destinationRoom = player.Room.East
	} else if strings.HasPrefix(dir, "w") {
		// west
		destinationRoom = player.Room.West
	} else if strings.HasPrefix(dir, "u") {
		// up
		destinationRoom = player.Room.Up
	} else if strings.HasPrefix(dir, "d") {
		// down
		destinationRoom = player.Room.Down
	} else {
		// unknown or unhandled direction
		log.Printf("Unknown move direction: %s", dir)
		return MoveResult{
			Success: false,
			Result:  "UNKNOWN_MOVE_DIRECTION",
		}
	}
	if move(destinationRoom, player) {
		return MoveResult{Success: true}
	} else {
		return MoveResult{
			Success: false,
			Result:  "CANT_MOVE",
		}
	}

}

func move(destinationRoom *world.Room, player *world.Player) bool {
	if destinationRoom != nil {
		destinationRoom.AddPlayer(player)
		return true
	}
	return false
}
