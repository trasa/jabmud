package world

import (
	"fmt"
)

type Player struct {
	Id            string
	Name          string
	Jid           string
	Room          *Room
	EventCallback func(*Player, interface{})
}

func NewPlayer(id string, jid string, name string, onEvent func(*Player, interface{})) *Player {
	p := Player{
		Name:          name,
		Jid:           jid,
		Id:            id,
		EventCallback: onEvent,
	}
	return &p
}

func (p *Player) String() string {
	return fmt.Sprintf("(Player Id='%s', Name='%s', Jid='%s' in room '%v')", p.Id, p.Name, p.Jid, p.Room)
}

func (p *Player) FindZone() *Zone {
	if p.Room != nil {
		return p.Room.Zone
	}
	return nil
}

func (p *Player) OnEvent(payload interface{}) {
	if p.EventCallback != nil {
		p.EventCallback(p, payload)
	}
}
