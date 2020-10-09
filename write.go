package main

import (
	"errors"
	"log"
)

func packetLength(d int) (byte, error) {
	if d >= -64 && d <= 63 {
		return 0x80, nil
	} else if d >= -8192 && d <= 8191 {
		return 0xa0, nil
	} else if d >= -1048576 && d <= 1048575 {
		return 0xc0, nil
	} else if d >= -134217728 && d <= 134217727 {
		return 0xe0, nil
	} else {
		return 0, errors.New("packetLength(): data out of range for dmm servo!")
	}
}

func dataBytes(d int) ([]byte, error) {

	var df byte = byte(((d & 0xFE00000) >> 21) | 0x80)
	var dh byte = byte(((d & 0x1FC000) >> 14) | 0x80)
	var dt byte = byte(((d & 0x3F80) >> 7) | 0x80)
	var do byte = byte((d & 0x7F) | 0x80)

	l, err := packetLength(d)
	if err != nil {
		log.Fatal(err)
	}

	if l == 0x80 {
		return []byte{do}, nil
	} else if l == 0xa0 {
		return []byte{do, dt}, nil
	} else if l == 0xc0 {
		return []byte{do, dt, dh}, nil
	} else if l == 0xe0 {
		return []byte{do, dt, dh, df}, nil
	} else {
		return nil, errors.New("dataBytes(): Could not parse data byte")
	}

}
