package main

import (
	"fmt"
	"log"
)

type Player struct {
	Id string
	Name string
	Jid  string
}
// go:generate stringer -type=Player

var knownPlayersById = make(map[string]Player)
var knownPlayersByJid = make(map[string]Player)

func FindPlayerByJid(jid string) Player {
	return knownPlayersByJid[jid]
}

func Login(player Player) error {
	if player.Id == "" {
		return fmt.Errorf("Login: JID '%s' didn't provide a valid Id", player.Jid)
	}
	log.Printf("Player '%s' (%s) (%s) logging in", player.Name, player.Id, player.Jid)
	knownPlayersById[player.Id] = player
	knownPlayersByJid[player.Jid] = player
	return nil
}

func Logout(player Player) error {
	if player.Id == "" {
		return fmt.Errorf("Logout: JID '%s' didn't provide a valid Id", player.Jid)
	}
	log.Printf("Player '%s' (%s) (%s) logged out", player.Name, player.Id, player.Jid)
	delete(knownPlayersById, player.Id)
	delete(knownPlayersByJid, player.Jid)
	return nil
}
