package world

import (
	"log"
	"testing"
)

func TestPlayerNotInZone(t *testing.T) {
	p := Player{}

	z := p.FindZone()
	if z != nil {
		t.Error("found a zone but should be nil")
	}

	r := p.FindRoom()
	if r != nil {
		t.Error("Found a room but should be nil")
	}
}

func TestPlayerInZoneButNotRoom(t *testing.T) {
	p := Player{ZoneId: startZone}

	z := p.FindZone()
	if z == nil {
		t.Errorf("Did not find zone but should have: %s", p.ZoneId)
	}
	log.Printf("Found zone: %v", z)

	r := p.FindRoom()
	if r != nil {
		t.Error("Found a room but should be nil")
	}
}

func TestPlayerInZoneAndRoom(t *testing.T) {
	p := Player{
		ZoneId: startZone,
		RoomId: startRoom,
	}

	z := p.FindZone()
	if z == nil {
		t.Errorf("Did not find zone but should have: %s", p.ZoneId)
	}
	log.Printf("Found zone: %v", z)

	r := p.FindRoom()
	if r == nil {
		t.Error("Did not find  room but should have: %s", p.RoomId)
	}
	log.Printf("Found room: %v", r)
}
