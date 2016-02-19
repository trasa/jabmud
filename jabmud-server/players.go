package main

import (
	"fmt"
	"log"
)

type Player struct {
	Id   string
	Name string
	Jid  string
}

func (p Player) String() string {
	return fmt.Sprintf("(Player Id='%s', Name='%s', Jid='%s'", p.Id, p.Name, p.Jid)
}

var knownPlayersById = make(map[string]Player)
var knownPlayersByJid = make(map[string]Player)

func FindPlayerByJid(jid string) Player {
	return knownPlayersByJid[jid]
}

func Login(player Player) error {
	if player.Id == "" {
		return fmt.Errorf("Login: JID '%s' didn't provide a valid Id", player.Jid)
	}
	log.Printf("%s logging in", player)
	knownPlayersById[player.Id] = player
	knownPlayersByJid[player.Jid] = player
	return nil
}

func Logout(player Player) error {
	if player.Id == "" {
		return fmt.Errorf("Logout: JID '%s' didn't provide a valid Id", player.Jid)
	}
	log.Printf("%s logged out", player)
	delete(knownPlayersById, player.Id)
	delete(knownPlayersByJid, player.Jid)
	return nil
}
