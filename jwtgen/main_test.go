package main

import (
	"strings"
	"testing"

	"garymenezes.com/xposer/test"
)

const (
	TESTCMD = "../bin/jwtgen"
)

func TestHelp(t *testing.T) {
	outStr, errStr, err := test.RunCommand(TESTCMD, []string{"-h"})
	if err == nil || err.Error() != "exit status 2" || !strings.Contains(errStr, "Usage of") {
		t.Error(outStr, errStr, err)
	}
}

func TestFlags(t *testing.T) {
	outStr, errStr, err := test.RunCommand(TESTCMD, nil)
	if err == nil || err.Error() != "exit status 1" || !strings.Contains(errStr, "you must enter a private key path") {
		t.Error(outStr, errStr, err)
	}

	outStr, errStr, err = test.RunCommand(TESTCMD, []string{"-d", "device"})
	if err == nil || err.Error() != "exit status 1" || !strings.Contains(errStr, "you must enter a private key path") {
		t.Error(outStr, errStr, err)
	}

	outStr, errStr, err = test.RunCommand(TESTCMD, []string{"-d", "device", "-o", "../bin/test.jwt"})
	if err == nil || err.Error() != "exit status 1" || !strings.Contains(errStr, "you must enter a private key path") {
		t.Error(outStr, errStr, err)
	}
}

func TestBadKeys(t *testing.T) {
	outStr, errStr, err := test.RunCommand(TESTCMD, []string{"-d", "device", "-o", "../bin/test.jwt", "-k", "../test/test.pub"})
	if err == nil || err.Error() != "exit status 1" {
		t.Error(outStr, errStr, err)
	}

	outStr, errStr, err = test.RunCommand(TESTCMD, []string{"-d", "device", "-o", "../bin/test.jwt", "-k", "main.go"})
	if err == nil || err.Error() != "exit status 1" || !strings.Contains(errStr, "Invalid Key") {
		t.Error(outStr, errStr, err)
	}
}

func TestSuccess(t *testing.T) {
	outStr, errStr, err := test.RunCommand(TESTCMD, []string{"-d", "device", "-o", "../bin/test.jwt", "-k", "../test/test.pem"})
	if err != nil {
		t.Error(outStr, errStr, err)
	}
}
