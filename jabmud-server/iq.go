package main

import (
	"github.com/emgee/go-xmpp/src/xmpp"
	"github.com/trasa/jabmud/commands"
	"github.com/trasa/jabmud/world"
	"log"
	"strings"
)

func HandleIq(iq *xmpp.Iq) *xmpp.Iq {
	log.Printf("iq: %T: %v", iq.Payload, iq.Payload)
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
	log.Printf("cmd: %s - %s", player, cmd)
	// so now go do something with the command...
	payload := commands.Serialize(commands.Run(player, cmd.Name, cmd.ArgList))
	response := iq.Response("result")
	response.Payload = payload
	log.Printf("sending response: %s", response.Payload)
	return response
}
