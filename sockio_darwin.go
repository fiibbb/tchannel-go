// +build darwin amd64

package tchannel

import "golang.org/x/sys/unix"

func getSendQueueLen(fd uintptr) (int, error) {
	// https://www.unix.com/man-page/osx/2/getsockopt/
	return unix.GetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_NWRITE)
}
