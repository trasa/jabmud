package world

import (
	"errors"
	"fmt"
)

type Room struct {
	Id          string
	Name        string
	Description string
	Zone        *Zone
	Players     map[string]*Player
	North       *Room
	South       *Room
	East        *Room
	West        *Room
	Up          *Room
	Down        *Room
}

func NewRoom(zone *Zone, id string, name string, description string, neighborRooms ...*Room) (result *Room, err error) {
	room := Room{
		Id:          id,
		Name:        name,
		Description: description,
		Zone:        zone,
		Players:     make(map[string]*Player),
	}
	for i, r := range neighborRooms {
		switch i {
		case 0:
			room.North = r
		case 1:
			room.South = r
		case 2:
			room.East = r
		case 3:
			room.West = r
		case 4:
			room.Up = r
		case 5:
			room.Down = r
		default:
			err = errors.New("Unknown room in neighborRooms")
		}
	}
	result = &room
	return
}

func (r Room) String() string {
	return fmt.Sprintf("(Room %s: '%s')", r.Id, r.Name)
}

func (r *Room) RemovePlayer(player *Player) {
	// TODO
}

func (r *Room) AddPlayer(player *Player) {
	r.Players[player.Id] = player
	player.Room = r
}
