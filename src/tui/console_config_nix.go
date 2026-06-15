//go:build !windows

package tui

import (
	"errors"
	"os"
	"syscall"
	"unsafe"
)

type winsize struct {
	Row, Col, Xpixel, Ypixel uint16
}

func fetchSize() (int, int, error) {
	var ws winsize
	_, _, err := syscall.Syscall(
		syscall.SYS_IOCTL,
		os.Stdout.Fd(),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&ws)),
	)
	if err != 0 {
		return 0, 0, errors.New("failed to get unix terminal info")
	}
	return int(ws.Col), int(ws.Row), nil
}
