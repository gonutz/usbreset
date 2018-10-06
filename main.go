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
		fmt.Fprintln(os.Stderr, "Usage: usbreset device-filename")
		return 1
	}
	filename := os.Args[1]

	fd, err := os.OpenFile(filename, os.O_WRONLY, 0)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	defer fd.Close()

	fmt.Println("Resetting USB device", filename)
	r1, r2, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		fd.Fd(),
		USBDEVFS_RESET,
		0,
	)
	fmt.Println(r1, r2, errno)

	return 0
}

/*
#include <stdio.h>
#include <unistd.h>
#include <fcntl.h>
#include <errno.h>
#include <sys/ioctl.h>

#include <linux/usbdevice_fs.h>


int main(int argc, char **argv)
{
	const char *filename;
	int fd;
	int rc;

	if (argc != 2) {
		fprintf(stderr, "Usage: usbreset device-filename\n");
		return 1;
	}
	filename = argv[1];

	fd = open(filename, O_WRONLY);
	if (fd < 0) {
		perror("Error opening output file");
		return 1;
	}

	printf("Resetting USB device %s\n", filename);
	rc = ioctl(fd, USBDEVFS_RESET, 0);
	if (rc < 0) {
		perror("Error in ioctl");
		return 1;
	}
	printf("Reset successful\n");

	close(fd);
	return 0;
}
*/
