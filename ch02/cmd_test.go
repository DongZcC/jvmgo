package main

import "testing"

func TestStarJvm(t *testing.T) {
	cmd := &Cmd{class: "java.lang.Object"}

	startJVM(cmd)
}
