package main

import (
	"fmt"
	"log"
)

type Player struct {
	Name string
	Jid  string
}

var knownPlayers = make(map[string]Player)

func Login(player Player) error {
	if player.Name == "" {
		return fmt.Errorf("Login: JID '%s' didn't provide a valid name", player.Jid)
	}
	log.Printf("Player '%s' (%s) logging in", player.Name, player.Jid)
	knownPlayers[player.Name] = player
	return nil
}

func Logout(player Player) error {
	if player.Name == "" {
		return fmt.Errorf("Logout: JID '%s' didn't provide a valid name", player.Jid)
	}
	log.Printf("Player '%s' (%s) logged out", player.Name, player.Jid)
	delete(knownPlayers, player.Name)
	return nil
}
