package tui

import (
	"errors"
	"os"
	"syscall"
	"unsafe"
)

var (
	kernel32                       = syscall.NewLazyDLL("kernel32.dll")
	procGetConsoleScreenBufferInfo = kernel32.NewProc("GetConsoleScreenBufferInfo")
)

type coord struct{ X, Y int16 }
type smallRect struct{ Left, Top, Right, Bottom int16 }
type consoleScreenBufferInfo struct {
	Size              coord
	CursorPosition    coord
	Attributes        uint16
	Window            smallRect
	MaximumWindowSize coord
}

func fetchSize() (int, int, error) {
	var csbi consoleScreenBufferInfo
	ret, _, _ := procGetConsoleScreenBufferInfo.Call(
		os.Stdout.Fd(),
		uintptr(unsafe.Pointer(&csbi)),
	)
	if ret == 0 {
		return 0, 0, errors.New("failed to get windows console info")
	}
	width := int(csbi.Window.Right - csbi.Window.Left + 1)
	height := int(csbi.Window.Bottom - csbi.Window.Top + 1)
	return width, height, nil
}
