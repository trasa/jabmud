package world

import (
	"fmt"
	"log"
)

var knownPlayersById = make(map[string]*Player)
var knownPlayersByJid = make(map[string]*Player)

func ClearKnownPlayers() {
	knownPlayersById = make(map[string]*Player)
	knownPlayersByJid = make(map[string]*Player)
}

func FindPlayerByJid(jid string) *Player {
	return knownPlayersByJid[jid]
}

func GetAllPlayers() (result []*Player) {
	for _, v := range knownPlayersById {
		result = append(result, v)
	}
	return result
}

func AddKnownPlayer(player *Player) {
	knownPlayersById[player.Id] = player
	knownPlayersByJid[player.Jid] = player
}

func Login(player *Player) error {
	if player.Id == "" {
		log.Printf("Player does not have a valid id: %s", player)
		return fmt.Errorf("Login: JID '%s' didn't provide a valid Id", player.Jid)
	}
	log.Printf("%s logging in", player)
	AddKnownPlayer(player)

	// put the player in the start room
	worldInstance.StartRoom.AddPlayer(player)
	return nil
}

func Logout(playerId string) error {
	log.Printf("%s logged out", playerId)
	player := knownPlayersById[playerId]
	player.Room.RemovePlayer(player)
	delete(knownPlayersById, playerId)
	if player != nil {
		delete(knownPlayersByJid, player.Jid)
	}
	return nil
}
