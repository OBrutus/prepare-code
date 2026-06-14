package system

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
)

// just put command we will make sure which shell to execute
func executeCommand(commands []string, targetDir string) (string, error) {
	shell := ""
	switch runtime.GOOS {
	case "windows":
		shell = "cmd"
	case "linux":
		shell = "bash"
	case "darwin":
		shell = "zsh"
	}

	if shell == "" {
		return "", errors.New("Unable to detect OS")
	}

	cmd := exec.Command(shell, commands...)
	if targetDir != "" {
		cmd.Dir = targetDir
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.New(fmt.Sprintf("Execution failed: %s", err))
	}

	return string(out), nil
}
