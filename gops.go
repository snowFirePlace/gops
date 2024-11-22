package gops

import (
	"bytes"
	"fmt"
	"os/exec"
)

// Output runs a given command using powershell and returns its output as an
// error and two strings. The first string is the standard output, and the
// second string is the standard error. The error is non-nil if the execution
// of the command fails.
func Output(command string) (string, error) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	cmd := exec.Command("powershell.exe", command)
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("%s", stderr.String())
	}
	return stdout.String(), err

}
