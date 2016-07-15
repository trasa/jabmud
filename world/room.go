package world

import "fmt"

type Room struct {
	Id          string
	Name        string
	Description string
	Zone        *Zone
	Players     []*Player
}

func (r Room) String() string {
	return fmt.Sprintf("(Room %s: '%s')", r.Id, r.Name)
}

func (r *Room) RemovePlayer(player *Player) {
	// TODO
}

func (r *Room) AddPlayer(player *Player) {
	r.Players = append(r.Players, player)
	player.Room = r
}
