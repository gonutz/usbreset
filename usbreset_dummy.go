//+build !linux

package usbreset

import "errors"

const USBDEVFS_RESET = 21780

// Reset only works on Linux.
func Reset(device string) error {
	return errors.New("usbreset only works on Linux")
}
