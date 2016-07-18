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

	r := NewRoom(&sampleZone, startRoomKey, "Central Portal", "It's a boring room, with boring stuff in it.")
	sampleZone.Rooms[r.Id] = r
	worldInstance.StartRoom = r

	// north room
	northRoom := NewRoom(&sampleZone, "northRoom", "North Room", "This room is north of the start.")
	sampleZone.Rooms[northRoom.Id] = northRoom

	r.North = northRoom
	northRoom.South = r

	log.Print("World built.")

}
