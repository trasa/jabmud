package world

import "testing"

func TestAddPlayerToRoom(t *testing.T) {
	room, err := NewRoom(nil, "id", "name", "description")
	if err != nil {
		panic(err)
	}
	player := Player{
		Id: "foo",
	}

	room.AddPlayer(&player)

	if player.Room != room {
		t.Error("player.Room is wrong")
	}
	if room.Players["foo"] == nil {
		t.Error("room.Players is wrong")
	}
}
