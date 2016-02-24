package world

import "fmt"

type Room struct {
	Id          string
	Name        string
	Description string
	Zone        Zone
}

func (r Room) String() string {
	return fmt.Sprintf("(Room %s: '%s')", r.Id, r.Name)
}
