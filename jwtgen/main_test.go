package main

import (
	"strings"
	"testing"

	"garymenezes.com/xfinity-xposer/test"
)

const (
	TESTCMD = "../bin/jwtgen"
)

func TestHelp(t *testing.T) {
	stdout, stderr, err := test.RunCommand(TESTCMD, []string{"-h"})
	if err == nil || err.Error() != "exit status 2" || !strings.Contains(stderr.String(), "Usage of") {
		t.Error(stdout.String(), stderr.String(), err)
	}
}

func TestFlags(t *testing.T) {
	stdout, stderr, err := test.RunCommand(TESTCMD, nil)
	if err == nil || err.Error() != "exit status 1" || !strings.Contains(stdout.String(), "you must enter a private key path") {
		t.Error(stdout.String(), stderr.String(), err)
	}

	stdout, stderr, err = test.RunCommand(TESTCMD, []string{"-d", "device"})
	if err == nil || err.Error() != "exit status 1" || !strings.Contains(stdout.String(), "you must enter a private key path") {
		t.Error(stdout.String(), stderr.String(), err)
	}

	stdout, stderr, err = test.RunCommand(TESTCMD, []string{"-d", "device", "-o", "../bin/test.jwt"})
	if err == nil || err.Error() != "exit status 1" || !strings.Contains(stdout.String(), "you must enter a private key path") {
		t.Error(stdout.String(), stderr.String(), err)
	}
}

func TestBadKeys(t *testing.T) {
	stdout, stderr, err := test.RunCommand(TESTCMD, []string{"-d", "device", "-o", "../bin/test.jwt", "-k", "../test/test.pub"})
	if err == nil || err.Error() != "exit status 1" {
		t.Error(stdout.String(), stderr.String(), err)
	}

	stdout, stderr, err = test.RunCommand(TESTCMD, []string{"-d", "device", "-o", "../bin/test.jwt", "-k", "main.go"})
	if err == nil || err.Error() != "exit status 1" || !strings.Contains(stdout.String(), "Invalid Key") {
		t.Error(stdout.String(), stderr.String(), err)
	}
}

func TestSuccess(t *testing.T) {
	stdout, stderr, err := test.RunCommand(TESTCMD, []string{"-d", "device", "-o", "../bin/test.jwt", "-k", "../test/test.pem"})
	if err != nil {
		t.Error(stdout.String(), stderr.String(), err)
	}
}
