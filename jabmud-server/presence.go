package main

import (
	"encoding/xml"
	"github.com/emgee/go-xmpp/src/xmpp"
	"github.com/trasa/jabmud/world"
	"log"
)

// Deal with an incoming Presence message, returning a presence
// message suitable to be returned to the client.
func HandlePresence(presence *xmpp.Presence) (response interface{}) {
	log.Printf("pres: type %s to %s from %s", presence.Type, presence.To, presence.From)
	// pres: type  to jabmud.localhost/mynick from tony.rasa@bw-mbp-trasa.glu.com/39118645601455771114633872
	tojid, _ := xmpp.ParseJID(presence.To)
	switch presence.Type {
	case "unavailable":
		// if type == unavailable then user has logged off
		log.Printf("logout name %s, jid %s", tojid.Resource, presence.From)
		world.Logout(tojid.Resource)

	case "":
		// create a new player object
		player := world.Player{Name: tojid.Resource, Jid: presence.From, Id: tojid.Resource}
		log.Printf("Attempting login for %s", player)
		if e := world.Login(&player); e != nil {
			log.Printf("Login failed for player %s", player)
			response = newErrorPresence(presence)
		} else {
			// success! response should reflect success case here...
			log.Printf("Login Success for player %s", player)
			response = newSuccessPresence(presence)
		}
	}
	return response
}

// SuccessPresence is a Presence that represents you successfully logging in.
//
//		<presence to='you@wherever/resource'
//				  from='jabmud.localhost/mynick'
//			<x xmlns='http://jabber.org/protocol/muc'>
//				<item affiliation='member' role='participant'></item>
//				<status code='110'></status>
//			</x>
//		</presence>
type SuccessPresence struct {
	xmpp.Presence
	X *MucPayload
}

// Create a new successful Presence response for this Presence request.
func newSuccessPresence(presence *xmpp.Presence) SuccessPresence {
	return SuccessPresence{
		Presence: xmpp.Presence{
			Id:   presence.Id,
			From: presence.To,
			To:   presence.From,
		},
		X: newSuccessMucPayload("member", "participant", "110"),
	}
}

// ErrorPresence is a Presence that contains an error.
//
// 		<presence to='you@wherever/resource'
//				  from='jabmud.localhost/mynick'>
//			<x xmlns='http://jabber.org/protocol/muc' />
//			<error type='cancel'>
//				<conflict xmlns='urn:ietf:params:xml:ns:xmpp-stanzas' />
//			</error>
//		</presence>
//
type ErrorPresence struct {
	xmpp.Presence
	X     *MucPayload
	Error ConflictError
}

// Build a new ErrorPresence based off of this Presence request.
func newErrorPresence(presence *xmpp.Presence) ErrorPresence {
	return ErrorPresence{
		Presence: xmpp.Presence{
			Id:   presence.Id,
			Type: "error",
			From: presence.To,
			To:   presence.From,
		},
		X: newEmptyMucPayload(),
		Error: ConflictError{
			Type:    "cancel",
			Payload: newConflictBlock(),
		},
	}
}

// MucPayload serializes to the MUC payload object.
//
//		<x xmlns='http://jabber.org/protocol/muc'>
//			<item affiliation='member' role='participant'></item>
//			<status code='110'></status>
// 		</x>
//
type MucPayload struct {
	XMLName xml.Name   `xml:"x"`
	Xmlns   string     `xml:"xmlns,attr"`
	Item    *MucItem   `xml:"item,omitempty"`
	Status  *MucStatus `xml:"status,omitempty"`
}

// Muc Item
//
//		<item affiliation='member' role='participant'></item>
type MucItem struct {
	Affiliation string `xml:"affiliation,attr,omitempty"`
	Role        string `xml:"role,attr,omitempty"`
}

// Muc Status Code
//
//		<status code='110'></status>
type MucStatus struct {
	Code string `xml:"code,attr,omitempty"`
}

// Creates an empty MUC payload object.
func newEmptyMucPayload() *MucPayload {
	return &MucPayload{Xmlns: "http://jabber.org/protocol/muc"}
}

// Creates a success Muc Payload for this affiliation, role, and code.
func newSuccessMucPayload(affiliation string, role string, code string) *MucPayload {
	mp := newEmptyMucPayload()
	mp.Item = &MucItem{
		Affiliation: affiliation,
		Role:        role,
	}
	if code != "" {
		mp.Status = &MucStatus{
			Code: code,
		}
	}
	return mp
}

// ConflictError is the inside payload of a Presence that results in an error.
//
//		<error type='cancel'>
//			<conflict xmlns='urn:ietf:params:xml:ns:xmpp-stanzas'></conflict>
//		</error>
type ConflictError struct {
	XMLName xml.Name       `xml:"error"`
	Type    string         `xml:"type,attr"`
	Payload *ConflictBlock `xml:",innerxml"`
}

// ConflictBlock is the inner conflict body of an error.
//
//		<conflict xmlns='urn:ietf:params:xml:ns:xmpp-stanzas'></conflict>
type ConflictBlock struct {
	XMLName xml.Name `xml:"conflict"`
	Xmlns   string   `xml:"xmlns,attr"`
}

// Create a new Conflict Block.
func newConflictBlock() *ConflictBlock {
	return &ConflictBlock{Xmlns: "urn:ietf:params:xml:ns:xmpp-stanzas"}
}
