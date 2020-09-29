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

func GetStringFromPipe(r io.Reader) (string, error) {
	bytes, err := ioutil.ReadAll(r)
	return string(bytes), err
}

func RunCommand(command string, args []string) (string, string, error) {
	stdout, stderr, err, cmd := StartCommand(command, args)
	outStr, err := GetStringFromPipe(stdout)
	if err != nil {
		return "", "", err
	}
	errStr, err := GetStringFromPipe(stderr)
	if err != nil {
		return outStr, errStr, err
	}
	return outStr, errStr, cmd.Wait()
}
