package jab

import (
	"github.com/emgee/go-xmpp/src/xmpp"
	"github.com/trasa/jabmud/commands"
	"github.com/trasa/jabmud/serde"
	"github.com/trasa/jabmud/world"
	"log"
	"strings"
)

func HandleIq(iq *xmpp.Iq) *xmpp.Iq {
	log.Printf("Handle IQ: %T: %v", iq.Payload, iq.Payload)
	if strings.HasPrefix(iq.Payload, "<command") {
		return handleIqCommand(iq)
	} else {
		log.Printf("Not a command-iq: %s", iq.Payload)
		return iq.Response("error")
	}
}

func handleIqCommand(iq *xmpp.Iq) *xmpp.Iq {
	cmd := DeserializeIqCommand(iq.Payload)
	player := world.FindPlayerByJid(iq.From)
	if player == nil {
		response := iq.Response("error")
		response.Payload = "Not Logged In"
		return response
	}

	// so now go do something with the command...
	payload := serde.Serialize(commands.Run(player, cmd.Name, cmd.ArgList))
	response := iq.Response("result")
	response.Payload = payload
	return response
}
