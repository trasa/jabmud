package main

import (
	"github.com/emgee/go-xmpp/src/xmpp"
	"log"
)

func main() {
	log.Printf("jabmud server started\n")
	log.Print("neat now with autocomplete working\n")

	jid, err := xmpp.ParseJID("jabmud.localhost")
	log.Printf("xmpp parsing: %v\n", err)
	stream, err := xmpp.NewStream("localhost:5275", nil)
	log.Printf("created stream: %v\n", err)
	X, err := xmpp.NewComponentXMPP(stream, jid, "secret")
	log.Printf("created component: %v\n", X)

	for i := range X.In {
		switch v := i.(type) {
		case error:
			log.Printf("error: %v\n", v)
		case *xmpp.Message:
			log.Printf("msg: %s says %s\n", v.From, v.Body)
			// for fun, send a response
			X.Out <- xmpp.Message{Body: "hi!", To: v.From, From: v.To}
		default:
			log.Printf("%T: %v\n", v, v)
		}
	}
}
