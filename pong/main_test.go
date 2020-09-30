package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
	"testing"
	"time"

	"garymenezes.com/xposer/test"
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
	if err == nil || err.Error() != "exit status 255" || !strings.Contains(errStr, "you must enter a public key path") {
		t.Error(outStr, errStr, err)
	}

	outStr, errStr, err = test.RunCommand(TESTCMD, []string{"--test"})
	if err == nil || err.Error() != "exit status 255" || !strings.Contains(errStr, "you must enter a public key path") {
		t.Error(outStr, errStr, err)
	}

	outStr, errStr, err = test.RunCommand(TESTCMD, []string{"-k", "../test/test.pem"})
	if err == nil || err.Error() != "exit status 255" || !strings.Contains(errStr, "couldn't read config") {
		t.Error(outStr, errStr, err)
	}
}

func TestBadKeys(t *testing.T) {
	outStr, errStr, err := test.RunCommand(TESTCMD, []string{"-c", "../test/test_conf.json", "-k", "../test/test.pem"})
	if err == nil || err.Error() != "exit status 255" {
		t.Error(outStr, errStr, err)
	}

	outStr, errStr, err = test.RunCommand(TESTCMD, []string{"-c", "../test/test_conf.json", "-k", "main.go"})
	if err == nil || err.Error() != "exit status 255" || !strings.Contains(errStr, "Invalid Key") {
		t.Error(outStr, errStr, err)
	}
}

func startServer(t *testing.T) (io.ReadCloser, io.ReadCloser, *exec.Cmd) {
	stdout, stderr, err, cmd := test.StartCommand(TESTCMD, []string{"--test", "-c", "../test/test_conf.json", "-k", "../test/test.pub"})
	if err != nil {
		outStr, _ := test.GetStringFromPipe(stdout)
		errStr, _ := test.GetStringFromPipe(stderr)
		t.Error(outStr, errStr, err)
	}

	// The test server should be running after this sleep
	time.Sleep(1 * time.Second)

	return stdout, stderr, cmd
}

func stopServer(t *testing.T, cmd *exec.Cmd) {
	// Stop the process
	if err := cmd.Process.Kill(); err != nil {
		t.Error("Couldn't kill test process: ", err)
	}
}

func TestPing(t *testing.T) {
	_, _, cmd := startServer(t)

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
	// Valid call
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

	stopServer(t, cmd)

	// Validate the database state
}
