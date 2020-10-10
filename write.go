package main

import (
	"errors"
)

var PacketLengthParseError = errors.New("packetLength(): data out of range for dmm servo.")
var DataParseError = errors.New("dataBytes(): data could not be parsed int 1-4 bytes.")

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
		return 0, PacketLengthParseError 
	}
}

func dataBytes(d int) ([]byte, error) {
	l, err := packetLength(d)
	if err != nil {
        return nil, PacketLengthParseError
	}

    // parse data into 4 bytes (7 bits long), and add start start bitr
	var df byte = byte(((d & 0xFE00000) >> 21) | 0x80)
	var dh byte = byte(((d & 0x1FC000) >> 14) | 0x80)
	var dt byte = byte(((d & 0x3F80) >> 7) | 0x80)
	var do byte = byte((d & 0x7F) | 0x80)

    //return byte slice deppending on packet length
	if l == 0x80 {
		return []byte{do}, nil
	} else if l == 0xa0 {
		return []byte{dt, do}, nil
	} else if l == 0xc0 {
		return []byte{dh, dt, do}, nil
	} else if l == 0xe0 {
		return []byte{df, dh, dt, do}, nil
	} else {
        return nil, DataParseError
    }

}
