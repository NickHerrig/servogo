package main

import (
	"errors"
)

var InvalidChecksumError = errors.New("validateChecksum(): The checksum packet was not valid.")

func validateChecksum(r []byte) (bool, error) {
	return true, nil
}

func ParsePacket(r []byte) (string, error) {
	//validateChecksum(r)
	return "motor response", nil
}
