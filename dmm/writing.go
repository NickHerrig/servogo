package dmm

import (
	"errors"
	"flag"
	"log"
)

// Errors returned from writing packet functions
var (
	PacketLengthParseError = errors.New("packetLength(): data out of range for dmm servo.")
	DataParseError         = errors.New("dataBytes(): data could not be parsed into 1-4 bytes.")
	FuncCodeNotImplemented = errors.New("funcCode(): That command isn't implemetnted.")
)

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

	// parse data into 4 bytes (7 bits long), and add start bit
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

/*
CreatePacket() creates a 4 - 7 long []byte from a drive ID
command, and data and returns it for writing to serial motor.

input validation is also handled through this function.

The packet follows the format:
  B0 - drive id byte
  B1 - packet length and function code byte
  B2 ~ B5 -  data bytes
  BN-1 - checksum byte

*/
func CreatePacket(id int, command string, data int) ([]byte, error) {

	// Validate user input, if error, print flag details and log failure
	err := validateInput(id, command, data)
	if err != nil {
		flag.PrintDefaults()
		log.Fatal(err)
	}

	// If command was passed with no data, change data to record in map[command]Functions.data
	// example: "servogo forwards" == servogo fowards --data 13000000
	if data == 0 {
		data = commandMap[command].data
	}

	var packet []byte

	// create the motor byte and append to packet
	mtid := byte(id)
	packet = append(packet, mtid)

	// create packetLength, FunctionCode, and append byte to packet
	pl, err := packetLength(data)
	if err != nil {
		return nil, err
	}
	fc, err := funcCode(command)
	if err != nil {
		return nil, err
	}
	plfcb := packetLengthFuncCodeByte(pl, fc)
	packet = append(packet, plfcb)

	// create databytes, iterate over []byte and append to packet
	dbs, err := dataBytes(data)
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
