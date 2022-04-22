package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmds := []*exec.Cmd{
		exec.Command("ps", "a"),
		exec.Command("grep", "signal"),
		exec.Command("grep", "-v", "grep"),
		exec.Command("grep", "-v", "go run"),
		exec.Command("awk", "{print $2}"),
	}
	fmt.Println(cmds)
}

func runCmds(cmds []*exec.Cmd) error {
	return nil
}
