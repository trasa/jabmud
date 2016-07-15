package world

import (
	"log"
	"testing"
)

func TestWorldHasStartRoom(t *testing.T) {
	log.Printf("Start Room: %v", worldInstance.StartRoom)
	if worldInstance.StartRoom == nil {
		t.Error("StartRoom should not be nil")
	}
}
