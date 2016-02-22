package world
import "fmt"

type Player struct {
	Id   string
	Name string
	Jid  string
}

func (p Player) String() string {
	return fmt.Sprintf("(Player Id='%s', Name='%s', Jid='%s')", p.Id, p.Name, p.Jid)
}

