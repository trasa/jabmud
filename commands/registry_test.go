package commands

import "testing"

func TestCommandRun(t *testing.T) {
	args := []string { "a", "b", "c"}
	Run("l", args)
	Run("look", nil)
	Run("notacommand", nil)
}
