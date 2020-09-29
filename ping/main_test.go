package main

import (
	"strings"
	"testing"

	"garymenezes.com/xfinity-xposer/test"
)

const (
	TESTCMD = "../bin/ping"
)

func TestHelp(t *testing.T) {
	outStr, errStr, err := test.RunCommand(TESTCMD, []string{"-h"})
	if err == nil || err.Error() != "exit status 2" || !strings.Contains(errStr, "Usage of") {
		t.Error(outStr, errStr, err)
	}
}

func TestFlags(t *testing.T) {
	outStr, errStr, err := test.RunCommand(TESTCMD, nil)
	if err == nil || err.Error() != "exit status 1" || !strings.Contains(errStr, "no such file or directory") {
		t.Error(outStr, errStr, err)
	}

	outStr, errStr, err = test.RunCommand(TESTCMD, []string{"-test"})
	if err == nil || err.Error() != "exit status 1" || !strings.Contains(errStr, "no such file or directory") {
		t.Error(outStr, errStr, err)
	}

	outStr, errStr, err = test.RunCommand(TESTCMD, []string{"-test", "-j", "../test/test.jwt"})
	if err != nil || !strings.Contains(errStr, "Failed to connect") {
		t.Error(outStr, errStr, err)
	}
}

func TestBadURL(t *testing.T) {
	outStr, errStr, err := test.RunCommand(TESTCMD, []string{"-test", "-j", "../test/test.jwt", "-u", "garbage"})
	if err != nil || !strings.Contains(errStr, "unsupported") {
		t.Error(outStr, errStr, err)
	}
}

func TestSuccess(t *testing.T) {
	outStr, errStr, err := test.RunCommand(TESTCMD, []string{"-test", "-j", "../test/test.jwt", "-u", "http://localhost:1234/ping"})
	if err != nil || !strings.Contains(errStr, "connection refused") {
		t.Error(outStr, errStr, err)
	}
}
