package commands

import "testing"

func TestCommandRun(t *testing.T) {
	Run("l")
	Run("look")
	Run("notacommand")
}