package jab

import (
	"github.com/emgee/go-xmpp/src/xmpp"
	"log"
)

var xmppcomponent *xmpp.XMPP

func Send(message interface{}) {
	xmppcomponent.Out <- message
}

func ConnectComponent(jabberHost string, jabberPort string) {
	// connect as component
	jid, _ := xmpp.ParseJID("jabmud.localhost")
	stream, _ := xmpp.NewStream(jabberHost+":"+jabberPort, nil)
	xmppcomponent, _ = xmpp.NewComponentXMPP(stream, jid, "secret")
	log.Printf("created component JID %v at %v\n", jid, xmppcomponent)

	for i := range xmppcomponent.In {
		switch v := i.(type) {
		case error:
			log.Printf("error: %v\n", v)

		case *xmpp.Message:
			log.Printf("msg: %s says %s\n", v.From, v.Body)
			// for fun, send a response
			Send(xmpp.Message{Body: "hi!", To: v.From, From: v.To, Type: "chat"})

		case *xmpp.Iq:
			if response := HandleIq(v); response != nil {
				log.Printf("Iq Response: %s", response)
				Send(response)
			}

		case *xmpp.Presence:
			// player name is in to:jabmud.localhost/(playername)
			if response := HandlePresence(v); response != nil {
				log.Printf("Presence Response: %s", response)
				Send(response)
			}

		default:
			log.Printf("(unhandled) %T: %v\n", v, v)
		}
	}
}
