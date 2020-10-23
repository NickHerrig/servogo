package main

import (
	"errors"
)

var InvalidChecksumError = errors.New("validateChecksum(): The checksum packet was not valid.")

func validateChecksum(p []byte) error {
	csb := p[len(p)-1] // checksum byte is last byte in packet
    bp := p[:len(p)-1]  // begining of the packet
    cs := checksumByte(bp)
    if csb == cs {
        return nil
    } else{
        return InvalidChecksumError
    }
}

func ParsePacket(p []byte) (string, error) {
	// Validate checksum byte from response
    err := validateChecksum(p)
	if err != nil {
        return "", err
    }
    /*TODO:
      Think through and design parsing of data...
      Packet length can be calculated with len(packet), is this needed? actually?
      need function to grab motor function code?
      parse data differently deppending on the function code?
    */

	return "parsed pkt", nil
}
