package world

import "fmt"

type Zone struct {
	Id    string
	Rooms map[string]*Room
	Name  string
}

func (z Zone) String() string {
	return fmt.Sprintf("(Zone %s: '%s')", z.Id, z.Name)
}
