package world

import "fmt"

type Player struct {
	Id     string
	Name   string
	Jid    string
	ZoneId string
	RoomId string
}

func (p *Player) String() string {
	return fmt.Sprintf("(Player Id='%s', Name='%s', Jid='%s')", p.Id, p.Name, p.Jid)
}

func (p *Player) FindZone() *Zone {
	return worldInstance.Zones[p.ZoneId]
}

// Find the Room this player is in.
func (p *Player) FindRoom() *Room {
	zone := p.FindZone()
	if zone != nil {
		return zone.Rooms[p.RoomId]
	}
	return nil
}
