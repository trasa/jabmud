package commands

import (
	"github.com/trasa/jabmud/serde"
	"github.com/trasa/jabmud/world"
	"log"
	"testing"
)

func TestWho(t *testing.T) {
	p := world.Player{
		Id:  "id",
		Jid: "jid",
	}
	world.AddKnownPlayer(&p)

	whoResult := Who(&p, nil).(WhoResult)
	log.Printf("whoResult %v", whoResult)
	if len(whoResult.PlayerIds) != 1 {
		t.Error("not enough players")
	}
	if whoResult.PlayerIds[0] != "id" {
		t.Error("wrong player id")
	}
}

func TestWhoResultXml(t *testing.T) {
	result := WhoResult{
		PlayerIds: []string{"a", "b", "c"},
	}
	log.Printf("result: %v", result)
	str := serde.Serialize(result)
	log.Printf("xml: %v", str)
}
