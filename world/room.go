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
	//for _, p := range r.Players {
	//	p.Tell(player has left room)
	//}
	player.Room = nil
}

func (r *Room) AddPlayer(player *Player) {
	if player.Room != nil {
		player.Room.RemovePlayer(player)
	}
	//for _, p := range r.Players {
	//	main.main.Send(main.NewSuccessPresence(p))
	//	p.Tell(player has entered room)
	//}
	r.Players[player.Id] = player
	player.Room = r
}
