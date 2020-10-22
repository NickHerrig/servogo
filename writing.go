package main

import (
	"errors"
	"log"
	"os"
	"strconv"
)

var PacketLengthParseError = errors.New("packetLength(): data out of range for dmm servo.")
var DataParseError = errors.New("dataBytes(): data could not be parsed int 1-4 bytes.")
var FuncCodeNotImplemented = errors.New("funcCode(): That command isn't implemetnted.")
var InvalidDriveIdError = errors.New("motorIdByte(): Drive Id must be 0 ~ 63")

func checksumByte(p []byte) byte {
	var cs byte
	for _, v := range p {
		cs += v
	}
	return (cs & 0x7f) | 0x80
}

func packetLengthFuncCodeByte(l, f byte) byte {
	return l | f
}

func funcCode(c string) (byte, error) {

	if _, ok := commandMap[c]; ok {
		return commandMap[c].code, nil
	} else {
		return 0, FuncCodeNotImplemented
	}
}

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

	// parse data into 4 bytes (7 bits long), and add start start bit
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

func motorIdByte(s string) (byte, error) {

	// Convert string env var to int or fail
	id, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	// Validate id is 0 ~ 63 or fail!
	if id < 0 || id > 63 {
		return 0, InvalidDriveIdError
	}

	return byte(id), nil
}

func CreatePacket(c string, d int) ([]byte, error) {
	/*
	   -------------------------------------------------------
	   | ID                          | One byte (Start byte) |
	   | packetLength + functioncode | One byte              |
	   | data                        | One to four bytes     |
	   | checksum                    | One byte              |
	   -------------------------------------------------------
	*/

	var packet []byte

	// fetch motor env var, create motor start byte, append to packet
	id, ok := os.LookupEnv("SERVO_DRIVE_ID")
	if !ok {
		log.Fatal("SERVO_DRIVE_ID env var not set")
	}
	motorId, err := motorIdByte(id)
	if err != nil {
		return nil, err
	}
	packet = append(packet, motorId)

	// create packetLength, FunctionCode, and append byte to packet
	pl, err := packetLength(d)
	if err != nil {
		return nil, err
	}
	fc, err := funcCode(c)
	if err != nil {
		return nil, err
	}
	plfcbyt := packetLengthFuncCodeByte(pl, fc)
	packet = append(packet, plfcbyt)

	// create databytes, iterate over []byte and append to packet
	dbs, err := dataBytes(d)
	if err != nil {
		return nil, err
	}
	for _, bt := range dbs {
		packet = append(packet, bt)
	}

	// create checksum byte from packet and append to packet
	cs := checksumByte(packet)
	packet = append(packet, cs)

	return packet, nil
}
