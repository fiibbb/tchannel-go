// +build linux amd64

package tchannel

import "golang.org/x/sys/unix"

func getSendQueueLen(fd uintptr) (int, error) {
	// https://linux.die.net/man/7/tcp
	return unix.IoctlGetInt(int(fd), unix.SIOCOUTQ)
}
