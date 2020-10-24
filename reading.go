package main

import (
	"errors"
)

var InvalidChecksumError = errors.New("validateChecksum(): The checksum packet was not valid.")

func validateChecksum(p []byte) error {
	csb := p[len(p)-1] // checksum byte is last byte in packet
	bp := p[:len(p)-1] // begining of the packet
	cs := checksumByte(bp)
	if csb == cs {
		return nil
	} else {
		return InvalidChecksumError
	}
}

func ParsePacket(p []byte) (string, error) {
	// Validate checksum byte from response
	err := validateChecksum(p)
	if err != nil {
		return "", err
	}
	// grab function code from byte two
	fc := p[1] | 0x1F

	// deppending on func code prase data differently
	if fc == 0x19 {
		//TODO: parse IS_STATUS return byte
	} else if fc == 0x1A {
		//TODO: parse IS_CONFIG return byte
	} else {
		//TODO: Parse normal data!
	}

	return "parsed pkt", nil
}
