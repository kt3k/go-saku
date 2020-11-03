// +build windows

package saku

import (
	"os"
	"os/exec"
	"strconv"
)

// Executes the given string as command
func execCommand(command string) *exec.Cmd {
	cmd := exec.Command("cmd.exe", "/s", "/c", command)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), "IN_SAKU=true")

	return cmd
}

// Terminates the given command
func terminateCommand(cmd *exec.Cmd) error {
	if cmd == nil || cmd.Process == nil {
		return nil
	}

	pid := cmd.Process.Pid
	c := exec.Command("taskkill", "/t", "/f", "/pid", strconv.Itoa(pid))
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	return c.Run()
}
