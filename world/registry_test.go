package world

import "testing"

func TestCommandRun(t *testing.T) {
	args := []string{"a", "b", "c"}
	p := Player { Id:"id", Jid:"jid", Name:"Guy"}
	Run(p, "l", args)
	Run(p, "look", nil)
	Run(p, "notacommand", nil)
}

func TestSerializeLookResult(t *testing.T) {
	lookResult := LookResult{Value: "You see nothing."}

	str := Serialize(lookResult)
	expected := "<LookResult><Value>You see nothing.</Value></LookResult>"
	if str != expected {
		t.Errorf("serialize didn't get expected string\nexp=%s\nact=%s", expected, str)
	}
}
