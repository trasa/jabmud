package main

import (
	"github.com/trasa/jabmud/commands"
	"fmt"
	"log"
)

var knownPlayersById = make(map[string]commands.Player)
var knownPlayersByJid = make(map[string]commands.Player)

func FindPlayerByJid(jid string) commands.Player {
	return knownPlayersByJid[jid]
}

func Login(player commands.Player) error {
	if player.Id == "" {
		return fmt.Errorf("Login: JID '%s' didn't provide a valid Id", player.Jid)
	}
	log.Printf("%s logging in", player)
	knownPlayersById[player.Id] = player
	knownPlayersByJid[player.Jid] = player
	return nil
}

func Logout(player commands.Player) error {
	if player.Id == "" {
		return fmt.Errorf("Logout: JID '%s' didn't provide a valid Id", player.Jid)
	}
	log.Printf("%s logged out", player)
	delete(knownPlayersById, player.Id)
	delete(knownPlayersByJid, player.Jid)
	return nil
}
