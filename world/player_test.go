package world

import (
	"log"
	"testing"
)

func TestPlayerNotInZone(t *testing.T) {
	p := Player{}

	r := p.Room
	if r != nil {
		t.Error("Found a room but should be nil")
	}

	z := p.FindZone()
	if z != nil {
		t.Error("found a zone but should be nil")
	}
}

func TestPlayerInZoneAndRoom(t *testing.T) {
	p := Player{
		Room: worldInstance.StartRoom,
	}

	z := p.FindZone()
	if z == nil {
		t.Errorf("Did not find zone but should have")
	}
	log.Printf("Found zone: %v", z)

	r := p.Room
	if r == nil {
		t.Error("Did not find  room but should have")
	}
	log.Printf("Found room: %v", r)
}
