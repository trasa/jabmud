package world

import "log"

type World struct {
	Zones     map[string]*Zone
	StartRoom *Room
}

var worldInstance World

var startZoneKey = "sample"
var startRoomKey = "start"

func init() {
	// Build a very boring world.
	worldInstance = World{Zones: make(map[string]*Zone)}
	sampleZone := Zone{
		Id:    startZoneKey,
		Name:  "Sample Zone",
		Rooms: make(map[string]*Room),
	}
	worldInstance.Zones[sampleZone.Id] = &sampleZone

	r := Room{
		Zone:        &sampleZone,
		Id:          startRoomKey,
		Name:        "Central Portal",
		Description: "It's a boring room, with boring stuff in it.",
	}
	sampleZone.Rooms[r.Id] = &r
	worldInstance.StartRoom = &r

	log.Print("World built.")

}
