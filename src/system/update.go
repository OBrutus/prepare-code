package system

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"prepare-code/src/config"
	"strings"
)

// TODO: Add auto update logic

var (
	ErrNonGitUpdateNotSupported error
)

const DefaultMessage = "Update Not supported, kindly reinstall!"

func Update() (string, error) {
	if !isGitInitiated() {
		return DefaultMessage, ErrNonGitUpdateNotSupported
	}

	return updateWithGit()
}

func isGitInitiated() bool {
	dir := config.GetInstallDir()

	info, err := os.Stat(filepath.Join(dir, ".git"))
	if err != nil {
		// Returns true if the error is specifically because the path does not exist
		if errors.Is(err, fs.ErrNotExist) {
			return false
		}
		// Any other error (e.g., permission denied) means we can't confirm
		// its existence, but treat it as false or handle it explicitly.
		return false
	}

	return info.IsDir()
}

func updateWithGit() (string, error) {
	var output string
	var err error

	// this considers that the repo is git initated
	output, err = executeCommand(
		strings.Split("git pull origin mainline", " "),
		config.GetInstallDir(),
	)
	if err != nil {
		return output, err
	}

	// now to install the update
	output, err = executeCommand(
		strings.Split("go build .", " "),
		config.GetInstallDir(),
	)

	return output, err
}
