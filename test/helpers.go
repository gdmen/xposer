package test

import (
	"io"
	"io/ioutil"
	"os/exec"
)

func StartCommand(command string, args []string) (io.ReadCloser, io.ReadCloser, error, *exec.Cmd) {
	cmd := exec.Command(command, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, err, cmd
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, nil, err, cmd
	}
	return stdout, stderr, cmd.Start(), cmd
}

func RunCommand(command string, args []string) (string, string, error) {
	stdout, stderr, err, cmd := StartCommand(command, args)
	bytesOut, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", "", err
	}
	bytesErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		return "", "", err
	}
	strOut := string(bytesOut)
	strErr := string(bytesErr)
	if err != nil {
		return strOut, strErr, err
	}
	return strOut, strErr, cmd.Wait()
}
