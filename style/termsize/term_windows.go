// +build windows

package termsize

import (
	"syscall"
	"unsafe"
)

type coord struct {
	x int16
	y int16
}

type smallRect struct {
	left   int16
	top    int16
	right  int16
	bottom int16
}

type consoleScreenBuffer struct {
	size       coord
	cursorPos  coord
	attrs      int32
	window     smallRect
	maxWinSize coord
}

func Width() (int, error) {
	hCon, err := syscall.Open("CONOUT$", syscall.O_RDONLY, 0)
	if err != nil {
		return -1, err
	}
	defer syscall.Close(hCon)

	sb, err := getConsoleScreenBufferInfo(hCon)
	if err != nil {
		return -1, err
	}
	return int(sb.size.x), nil
}

func getConsoleScreenBufferInfo(hCon syscall.Handle) (sb consoleScreenBuffer, err error) {
	modkernel32 := syscall.NewLazyDLL("kernel32.dll")
	procGetConScrBufInfo := modkernel32.NewProc("GetConsoleScreenBufferInfo")

	rc, _, errno := syscall.Syscall(procGetConScrBufInfo.Addr(), 2, uintptr(hCon), uintptr(unsafe.Pointer(&sb)), 0)
	if rc == 0 {
		err = syscall.Errno(errno)
	}
	return
}
