package main

import "testing"

func TestExecutor(t *testing.T) {
	Run("l")
	Run("look")
	Run("notacommand")
}
