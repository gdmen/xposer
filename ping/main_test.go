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
	stdout, stderr, err := test.RunCommand(TESTCMD, []string{"-h"})
	if err == nil || err.Error() != "exit status 2" || !strings.Contains(stderr, "Usage of") {
		t.Error(stdout, stderr, err)
	}
}

func TestFlags(t *testing.T) {
	stdout, stderr, err := test.RunCommand(TESTCMD, nil)
	if err == nil || err.Error() != "exit status 1" || !strings.Contains(stderr, "no such file or directory") {
		t.Error(stdout, stderr, err)
	}

	stdout, stderr, err = test.RunCommand(TESTCMD, []string{"-test"})
	if err == nil || err.Error() != "exit status 1" || !strings.Contains(stderr, "no such file or directory") {
		t.Error(stdout, stderr, err)
	}

	stdout, stderr, err = test.RunCommand(TESTCMD, []string{"-test", "-j", "../test/test.jwt"})
	if err != nil || !strings.Contains(stderr, "Failed to connect") {
		t.Error(stdout, stderr, err)
	}
}

func TestBadURL(t *testing.T) {
	stdout, stderr, err := test.RunCommand(TESTCMD, []string{"-test", "-j", "../test/test.jwt", "-u", "garbage"})
	if err != nil || !strings.Contains(stderr, "unsupported") {
		t.Error(stdout, stderr, err)
	}
}

func TestSuccess(t *testing.T) {
	stdout, stderr, err := test.RunCommand(TESTCMD, []string{"-test", "-j", "../test/test.jwt", "-u", "http://localhost:1234/ping"})
	if err != nil || !strings.Contains(stderr, "connection refused") {
		t.Error(stdout, stderr, err)
	}
}
