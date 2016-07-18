package commands

import (
	"testing"
	"log"
)

func TestWhoXml(t *testing.T) {
	result := WhoResult{
		PlayerIds: []string{"a", "b", "c"},
	}
	log.Printf("result: %v", result)
	str := Serialize(result)
	log.Printf("xml: %v", str)
}