package main

import (
	"github.com/emgee/go-xmpp/src/xmpp"
	"github.com/trasa/jabmud/commands"
	"log"
	"strings"
)

func main() {
	go connectHttpServer()
	connectComponent()
}

func connectComponent() {
	// connect as component
	jid, _ := xmpp.ParseJID("jabmud.localhost")
	stream, _ := xmpp.NewStream("localhost:5275", nil)
	X, _ := xmpp.NewComponentXMPP(stream, jid, "secret")
	log.Printf("created component JID %v at %v\n", jid, X)

	for i := range X.In {
		switch v := i.(type) {
		case error:
			log.Printf("error: %v\n", v)

		case *xmpp.Message:
			log.Printf("msg: %s says %s\n", v.From, v.Body)
			// for fun, send a response
			X.Out <- xmpp.Message{Body: "hi!", To: v.From, From: v.To, Type: "chat"}

		case *xmpp.Iq:
			log.Printf("iq: %T: %v", v.Payload, v.Payload)
			if strings.HasPrefix(v.Payload, "<command") {
				cmd := ParseCommand(v.Payload)
				log.Printf("cmd: %s", cmd)
				// so now go do something with the command...
				commands.Run(cmd.Name, cmd.ArgList)
				response := v.Response("result")
				response.Payload = "neat"
				X.Out <- response
			} else {
				log.Printf("Not a command-iq: %s", v.Payload)
				// now what?
			}

		default:
			log.Printf("%T: %v\n", v, v)
		}
	}
}
