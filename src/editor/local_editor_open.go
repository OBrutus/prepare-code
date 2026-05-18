package editor

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"prepare-code/src/config"
)

var (
	ErrOpenNotPreffered = errors.New("User preffered not to open in editor")
)

func TryOpenInEditor(fileName string) error {
	editor, toOpen := config.GetEditor()
	if !toOpen {
		return ErrOpenNotPreffered
	}

	fmt.Println("Opening the file:", fileName, "in", editor)
	cmd := exec.Command(editor, fileName)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
