package test

import (
	"bytes"
	"os/exec"
)

func RunCommand(command string, args []string) (stdout, stderr bytes.Buffer, err error) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	return stdout, stderr, cmd.Run()
}
