package utils

import (
	"bytes"
	"os/exec"
	"strconv"
	"strings"
)

func ExecuteCmd(command string, args ...string) (error) {
	cmd := exec.Command(command, args...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		_, ok := err.(*exec.ExitError)
		if !ok {
			return err
		}	
	}


	return nil
}

func UnquoteCodePoint(s string) (string) {
    r, err := strconv.ParseInt(strings.TrimPrefix(s, "\\U"), 16, 32)
	if err != nil {
		return "(\\|) ._. (|/)"
	}
    return string(r)
}