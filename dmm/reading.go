package dmm 

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

// Error codes returned by reading packet from servo
var (
	InvalidChecksumError = errors.New("validateChecksum(): The checksum packet was not valid.")
)

type motorStatus struct {
	motion string
	alarm  string
}

func parseStatus(d byte) string {
	s := motorStatus{}

	bit := byte(d)
	alarmBit := (bit >> 2) & 0b111

	if alarmBit == 0 {
		s.alarm = "no alarm"
	} else if alarmBit == 1 {
		s.alarm = "lost phase"
	} else if alarmBit == 2 {
		s.alarm = "over current"
	} else if alarmBit == 3 {
		s.alarm = "overheat/overpower"
	} else if alarmBit == 4 {
		s.alarm = "rcr error"
	} else {
		s.alarm = "TBD"
	}

	motionBit := (d >> 5) & 1
	if motionBit == 0 {
		s.motion = "completed"
	} else {
		s.motion = "busy"
	}

	st, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	msg := string(st)

	return msg
}

func parseData(d []byte) string {
	p := int(d[0] & 0x7f)
	p = (p & (0x40 - 1)) - (p & 0x40)
	for _, b := range d[1:] {
		p = (p << 7) + int(b&0x7f)
	}
	msg := fmt.Sprint(p)

	return msg
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

func ParsePacket(p []byte) (string, error) {
	// Validate checksum byte from response
	err := validateChecksum(p)
	if err != nil {
		return "", err
	}

	fc := p[1] | 0x1F    // grab function code from byte two
	d := p[2 : len(p)-1] // grab data bytes

	var msg string

	// deppending on func code parse data differently
	if fc == 0x19 {
		msg = parseStatus(d[0])
	} else if fc == 0x1A {
		msg = "not implemented"
	} else {
		msg = parseData(d)
	}

	return msg, nil
}
