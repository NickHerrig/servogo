package main

import (
	"errors"
)

func packetLength(d int) (byte, error) {
	if (d >= -64) && (d <= 63) {
		return 0x80, nil
	} else if (d >= -8192) && (d <= 8191) {
		return 0xa0, nil
	} else if (d >= -1048576) && (d <= 1048575) {
		return 0xc0, nil
	} else if (d >= -134217728) && (d <= 134217727) {
		return 0xe0, nil
	} else {
		return 0, errors.New("data out of range for dmm servo!")
	}
}
