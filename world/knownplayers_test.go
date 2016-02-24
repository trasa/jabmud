package world

import (
	"log"
	"testing"
)

func TestLogin(t *testing.T) {
	ClearKnownPlayers()
	originalPlayer := Player{Id: "id", Name: "Name", Jid: "Jid"}
	if err := Login(originalPlayer); err != nil {
		t.Error("Failed to login")
	}

	p := FindPlayerByJid("Jid")
	if originalPlayer != p {
		t.Errorf("Found different player than expected: expected %s, actual %s", originalPlayer, p)
	}
}

func TestGetAll(t *testing.T) {
	ClearKnownPlayers()
	playerA := Player{"a", "a", "a"}
	playerB := Player{"b", "b", "b"}
	if err := Login(playerA); err != nil {
		t.Error("Failed to login A")
	}
	if err := Login(playerB); err != nil {
		t.Error("Failed to login B")
	}

	players := GetAllPlayers()
	if len(players) != 2 {
		t.Errorf("Expected 2 players, found %d", len(players))
	}
	for _, p := range players {
		log.Printf("Player: %s", p)
	}
}
