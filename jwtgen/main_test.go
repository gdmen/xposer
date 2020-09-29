package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func RunCommand(command string, args []string) (stdout, stderr bytes.Buffer, err error) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	return stdout, stderr, cmd.Run()
}

func TestHelp(t *testing.T) {
	stdout, stderr, err := RunCommand("../bin/jwtgen", []string{"-h"})
	if err == nil || err.Error() != "exit status 2" || !strings.Contains(stderr.String(), "Usage of") {
		t.Error(stdout.String(), stderr.String(), err)
	}
}

func TestMissingArgs(t *testing.T) {
	stdout, stderr, err := RunCommand("../bin/jwtgen", nil)
	if err == nil || err.Error() != "exit status 1" || !strings.Contains(stdout.String(), "you must enter a private key path") {
		t.Error(stdout.String(), stderr.String(), err)
	}

	stdout, stderr, err = RunCommand("../bin/jwtgen", []string{"-d", "device"})
	if err == nil || err.Error() != "exit status 1" || !strings.Contains(stdout.String(), "you must enter a private key path") {
		t.Error(stdout.String(), stderr.String(), err)
	}

	stdout, stderr, err = RunCommand("../bin/jwtgen", []string{"-d", "device", "-o", "../bin/test.jwt"})
	if err == nil || err.Error() != "exit status 1" || !strings.Contains(stdout.String(), "you must enter a private key path") {
		t.Error(stdout.String(), stderr.String(), err)
	}
}

func TestBadKeys(t *testing.T) {
	stdout, stderr, err := RunCommand("../bin/jwtgen", []string{"-d", "device", "-o", "../bin/test.jwt", "-k", "../test/test.pub"})
	if err == nil || err.Error() != "exit status 1" {
		t.Error(stdout.String(), stderr.String(), err)
	}

	stdout, stderr, err = RunCommand("../bin/jwtgen", []string{"-d", "device", "-o", "../bin/test.jwt", "-k", "main.go"})
	if err == nil || err.Error() != "exit status 1" || !strings.Contains(stdout.String(), "Invalid Key") {
		t.Error(stdout.String(), stderr.String(), err)
	}
}

func TestSuccess(t *testing.T) {
	stdout, stderr, err := RunCommand("../bin/jwtgen", []string{"-d", "device", "-o", "../bin/test.jwt", "-k", "../test/test.pem"})
	if err != nil {
		t.Error(stdout.String(), stderr.String(), err)
	}
}
