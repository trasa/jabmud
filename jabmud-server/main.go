package main

import (
	"github.com/emgee/go-xmpp/src/xmpp"
	"log"
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
			cmd := ParseCommand(v.Payload)
			log.Printf("cmd: %s", cmd)

		default:
			log.Printf("%T: %v\n", v, v)
		}
	}
}
