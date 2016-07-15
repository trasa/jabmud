package world

import "fmt"

type Player struct {
	Id   string
	Name string
	Jid  string
	Room *Room
}

func (p *Player) String() string {
	return fmt.Sprintf("(Player Id='%s', Name='%s', Jid='%s')", p.Id, p.Name, p.Jid)
}

func (p *Player) FindZone() *Zone {
	if p.Room != nil {
		return p.Room.Zone
	}
	return nil
}
