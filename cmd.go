package common

import "os/exec"

// CommandExists checks if a command exists.
func CommandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
