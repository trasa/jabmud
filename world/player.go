package world

import (
	"fmt"
	"log"
)

type Player struct {
	Id           string
	Name         string
	Jid          string
	Room         *Room
	EventChannel chan interface{}
}

func NewPlayer(id string, jid string, name string) *Player {
	p := Player{
		Name:         name,
		Jid:          jid,
		Id:           id,
		EventChannel: make(chan interface{}),
	}
	go func() {
		log.Printf("begin recieve")
		rec := <-p.EventChannel
		log.Printf("event channel recieved %v", rec)
	}()
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
