package main

import (
	"errors"
)

var InvalidChecksumError = errors.New("validateChecksum(): The checksum packet was not valid.")

func parseData(d []byte) int {
	p := int(d[0] & 0x7f)
	p = (p & (0x40 - 1)) - (p & 0x40)
	for _, b := range d[1:] {
		p = (p << 7) + int(b&0x7f)
	}

	return p
}

func validateChecksum(p []byte) error {
	csb := p[len(p)-1] // checksum byte is last byte in packet
	bp := p[:len(p)-1] // slice of the rest of the packet
	cs := checksumByte(bp)
	if csb == cs {
		return nil
	} else {
		return InvalidChecksumError
	}
}

func ParsePacket(p []byte) (int, error) {
	// Validate checksum byte from response
	err := validateChecksum(p)
	if err != nil {
		return 0, err
	}

	fc := p[1] | 0x1F    // grab function code from byte two
	d := p[2 : len(p)-1] // grab data bytes

	var msg int
	// deppending on func code prase data differently
	if fc == 0x19 {
		//TODO: parse IS_STATUS return byte
	} else if fc == 0x1A {
		//TODO: parse IS_CONFIG return byte
	} else {
		msg = parseData(d)
	}

	return msg, nil
}
