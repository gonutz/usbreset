//+build linux

// usbreset -- send a USB port reset to a USB device
package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	os.Exit(runMain())
}

func runMain() int {
	const USBDEVFS_RESET = 21780

	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, `Usage: usbreset device-filename
  File names look like this:
    /dev/bus/usb/xxx/xxx
  where the xxx are replaced by three digits.
`)
		return 1
	}
	filename := os.Args[1]

	fd, err := os.OpenFile(filename, os.O_WRONLY, 0)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening output file:", err)
		return 1
	}
	defer fd.Close()

	fmt.Println("Resetting USB device", filename)
	ret, _, err := syscall.Syscall(
		syscall.SYS_IOCTL,
		fd.Fd(),
		USBDEVFS_RESET,
		0,
	)
	if ret != 0 {
		fmt.Fprintln(os.Stderr, "Error in ioctl:", err)
		return 1
	}
	fmt.Println("Reset successful")

	return 0
}
