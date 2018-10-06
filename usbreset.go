//+build linux

package usbreset

import (
	"errors"
	"fmt"
	"os"
	"syscall"
)

const USBDEVFS_RESET = 21780

// Reset resets the given USB device, you typically need elevated rights to do
// this. A device name usually looks like this: /dev/bus/usb/xxx/xxx where the
// xxx are replaced by three digits.
func Reset(device string) error {
	fd, err := os.OpenFile(device, os.O_WRONLY, 0)
	if err != nil {
		return errors.New("Error opening output file: " + err.Error())
	}
	defer fd.Close()

	ret, _, err := syscall.Syscall(
		syscall.SYS_IOCTL,
		fd.Fd(),
		USBDEVFS_RESET,
		0,
	)
	if ret != 0 {
		return errors.New("Error in ioctl: " + err.Error())
	}

	return nil
}
