package world

import (
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

func NewRoom(zone *Zone, id string, name string, description string) *Room {
	return &Room{
		Id:          id,
		Name:        name,
		Description: description,
		Zone:        zone,
		Players:     make(map[string]*Player),
	}
}

func (r Room) String() string {
	return fmt.Sprintf("(Room %s: '%s')", r.Id, r.Name)
}

func (r *Room) RemovePlayer(player *Player) {
	delete(r.Players, player.Id)
	// tell players in room that this player has left
	for _, p := range r.Players {
		p.OnEvent(ExitRoomEvent{
			PlayerId: player.Id,
			ZoneId:   r.Zone.Id,
			RoomId:   r.Id,
		})
	}
	player.Room = nil
}

func (r *Room) AddPlayer(player *Player) {
	if player.Room != nil {
		player.Room.RemovePlayer(player)
	}
	for _, p := range r.Players {
		p.OnEvent(EnterRoomEvent{
			PlayerId: player.Id,
			ZoneId:   r.Zone.Id,
			RoomId:   r.Id,
		})
	}
	r.Players[player.Id] = player
	player.Room = r
}

type ExitRoomEvent struct {
	PlayerId string
	ZoneId   string
	RoomId   string
}

type EnterRoomEvent struct {
	PlayerId string
	ZoneId   string
	RoomId   string
}
