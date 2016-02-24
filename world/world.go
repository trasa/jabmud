package world

import "log"

type World struct {
	Zones map[string]*Zone
}

var worldInstance World

var startZone = "sample"
var startRoom = "start"

func init() {
	// Build a very boring world.
	worldInstance = World{Zones: make(map[string]*Zone)}
	sampleZone := Zone{
		Id:    startZone,
		Name:  "Sample Zone",
		Rooms: make(map[string]*Room),
	}
	worldInstance.Zones[sampleZone.Id] = &sampleZone

	r := Room{
		Zone:        sampleZone,
		Id:          startRoom,
		Name:        "Central Portal",
		Description: "It's a boring room, with boring stuff in it.",
	}
	sampleZone.Rooms[r.Id] = &r

	log.Print("World built.")

}

func StartRoom() *Room {
	return worldInstance.Zones[startZone].Rooms[startRoom]
}
