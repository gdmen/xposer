package main

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"garymenezes.com/xfinity-xposer/test"
)

const (
	TESTCMD = "../bin/test_pong"
)

func TestHelp(t *testing.T) {
	outStr, errStr, err := test.RunCommand(TESTCMD, []string{"-h"})
	if err == nil || err.Error() != "exit status 2" || !strings.Contains(errStr, "Usage of") {
		t.Error(outStr, errStr, err)
	}
}

func TestFlags(t *testing.T) {
	outStr, errStr, err := test.RunCommand(TESTCMD, nil)
	if err == nil || err.Error() != "exit status 1" || !strings.Contains(errStr, "you must enter a public key path") {
		t.Error(outStr, errStr, err)
	}
}

func TestBadKeys(t *testing.T) {
	outStr, errStr, err := test.RunCommand(TESTCMD, []string{"-k", "../test/test.pem"})
	if err == nil || err.Error() != "exit status 1" {
		t.Error(outStr, errStr, err)
	}

	outStr, errStr, err = test.RunCommand(TESTCMD, []string{"-k", "main.go"})
	if err == nil || err.Error() != "exit status 1" || !strings.Contains(errStr, "Invalid Key") {
		t.Error(outStr, errStr, err)
	}
}

func TestSuccess(t *testing.T) {
	stdout, stderr, err, cmd := test.StartCommand(TESTCMD, []string{"-k", "../test/test.pub"})
	var outStr, errStr string
	if err != nil {
		outStr, _ = test.GetStringFromPipe(stdout)
		errStr, _ = test.GetStringFromPipe(stderr)
		t.Error(outStr, errStr, err)
	}

	// The test server should be running after this sleep
	time.Sleep(1 * time.Second)

	// Make some test HTTP requests
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/ping", nil)
	if err != nil {
		t.Error(err)
	}

	// No Authorization Header
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	} else {
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusBadRequest {
			body, _ := ioutil.ReadAll(resp.Body)
			t.Error(string(body), resp.StatusCode)
		}
	}

	// Empty Authorization Header
	req.Header.Add("Authorization", "")
	resp, err = client.Do(req)
	if err != nil {
		t.Error(err)
	} else {
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusBadRequest {
			body, _ := ioutil.ReadAll(resp.Body)
			t.Error(string(body), resp.StatusCode)
		}
	}

	// Nonsense Authorization Header
	req.Header.Set("Authorization", "gobbledygook")
	resp, err = client.Do(req)
	if err != nil {
		t.Error(err)
	} else {
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusBadRequest {
			body, _ := ioutil.ReadAll(resp.Body)
			t.Error(string(body), resp.StatusCode)
		}
	}

	// Malformed Authorization Header
	req.Header.Set("Authorization", "Bearer: jwt")
	resp, err = client.Do(req)
	if err != nil {
		t.Error(err)
	} else {
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusBadRequest {
			body, _ := ioutil.ReadAll(resp.Body)
			t.Error(string(body), resp.StatusCode)
		}
	}

	// Bad JWT
	req.Header.Set("Authorization", "Bearer not.real.jwt")
	resp, err = client.Do(req)
	if err != nil {
		t.Error(err)
	} else {
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusUnauthorized {
			body, _ := ioutil.ReadAll(resp.Body)
			t.Error(string(body), resp.StatusCode)
		}
	}

	jwt, err := ioutil.ReadFile("../test/test.jwt")
	if err != nil {
		t.Error(err)
	}
	// Success
	req.Header.Set("Authorization", "Bearer "+string(jwt))
	resp, err = client.Do(req)
	if err != nil {
		t.Error(err)
	} else {
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			body, _ := ioutil.ReadAll(resp.Body)
			t.Error(string(body), resp.StatusCode)
		}
	}

	// Stop the process
	if err = cmd.Process.Kill(); err != nil {
		t.Error("Couldn't kill test process: ", err)
	}
}
