// +build windows

package text

import (
	"os"

	"golang.org/x/sys/windows"
)

func isANSISupported() bool {
	outHandle := windows.Handle(os.Stdout.Fd())
	var outMode uint32
	if err := windows.GetConsoleMode(outHandle, &outMode); err == nil {
		if outMode&windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING != 0 {
			return true
		}
		if err := windows.SetConsoleMode(outHandle, outMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING); err == nil {
			return true
		}
	}
	return false
}
