package jab

import (
	"github.com/emgee/go-xmpp/src/xmpp"
	"github.com/trasa/jabmud/world"
	"log"
)

// Event Handler for sending messages to a player
func OnPlayerEvent(player *world.Player, payload interface{}) {
	log.Printf("Sending Player Event to player %v: %v", player, payload)
	Send(xmpp.Message{Body: Serialize(payload), To: player.Jid, From: "jabmud.localhost", Type: "chat"})
}
