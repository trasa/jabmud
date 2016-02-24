package world

import "log"

type World struct {
	Zones map[string]Zone
}

var worldInstance World

func init() {
	// Build a very boring world.
	worldInstance = World{Zones: make(map[string]Zone)}
	sampleZone := Zone{
		Id:    "Sample",
		Name:  "Sample",
		Rooms: make(map[string]Room),
	}
	worldInstance.Zones[sampleZone.Id] = sampleZone

	r := Room{Id: "start", Name: "Central Portal"}
	sampleZone.Rooms[r.Id] = r

	log.Print("World built.")

}
