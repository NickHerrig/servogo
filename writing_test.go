package main

import (
	"bytes"
	"errors"
	"strconv"
	"testing"
)

func TestPacketLength(t *testing.T) {
	tests := []struct {
		name      string
		input     int
		want      byte
		wantError error
	}{
		{"one positive data packet", 50, 0x80, nil},
		{"one negative data packet", -50, 0x80, nil},
		{"two positive data packets", 8000, 0xA0, nil},
		{"two negative data packets", -8000, 0xA0, nil},
		{"three positive data packets", 1000000, 0xC0, nil},
		{"three negative data packets", -1000000, 0xC0, nil},
		{"four positive data packets", 130000000, 0xE0, nil},
		{"four negative data packets", -130000000, 0xE0, nil},
		{"five positive data packets", 140000000, 0, PacketLengthParseError},
		{"five negative data packets", -140000000, 0, PacketLengthParseError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pl, err := packetLength(tt.input)

			if err != nil {
				if errors.Is(err, tt.wantError) != true {
					t.Errorf("want %q; got %q", tt.wantError, err)
				}
			} else {
				if pl != tt.want {
					t.Errorf("want %q; got %q", tt.want, pl)
				}
				if err != tt.wantError {
					t.Errorf("want %q; got %q", tt.wantError, err)
				}
			}
		})
	}
}

func TestDataBytes(t *testing.T) {
	tests := []struct {
		name      string
		input     int
		want      []byte
		wantError error
	}{
		{"one positive data packet", 60, []byte{0xBC}, nil},
		{"one negative data packet", -60, []byte{0xC4}, nil},
		{"two positive data packet", 150, []byte{0x81, 0x96}, nil},
		{"two negative data packet", -150, []byte{0xFE, 0xEA}, nil},
		{"three positive data packet", 9000, []byte{0x80, 0xC6, 0xA8}, nil},
		{"three negative data packet", -9000, []byte{0xFF, 0xB9, 0xD8}, nil},
		{"four positive data packet", 120000000, []byte{0xB9, 0x9C, 0x9C, 0x80}, nil},
		{"four negative data packet", -120000000, []byte{0xC6, 0xE3, 0xE4, 0x80}, nil},
		{"error five positive data packet", 140700900, nil, PacketLengthParseError},
		{"error five negative data packet", -140080900, nil, PacketLengthParseError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := dataBytes(tt.input)

			if err != nil {
				if errors.Is(err, tt.wantError) != true {
					t.Errorf("want %q; got %q", tt.wantError, err)
				}
			} else {
				if bytes.Equal(b, tt.want) != true {
					t.Errorf("want %q; got %q", tt.want, b)
				}
				if err != tt.wantError {
					t.Errorf("want %q; got %q", tt.wantError, err)
				}
			}
		})
	}
}

func TestChecksumByte(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  byte
	}{
		{"four byte packet", []byte{0x02, 0x83, 0x80}, 0x85},
		{"five byte packet", []byte{0x03, 0x82, 0x82, 0x86}, 0x8d},
		{"six byte packet", []byte{0x01, 0x81, 0x81, 0x81, 0x81}, 0x85},
		{"seven byte packet", []byte{0x04, 0x81, 0x81, 0x81, 0x81, 0x81}, 0x89},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := checksumByte(tt.input)
			if cs != tt.want {
				t.Errorf("want %q; got %q", tt.want, cs)
			}
		})
	}
}

func TestPacketLengthFuncCodeByte(t *testing.T) {
	tests := []struct {
		name      string
		packetLen byte
		funcCode  byte
		want      byte
	}{
		{"four packet, relative func code", 0x80, 0x03, 0x83},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := packetLengthFuncCodeByte(tt.packetLen, tt.funcCode)
			if b != tt.want {
				t.Errorf("want %q; got %q", tt.want, b)
			}
		})
	}
}

func TestMotorIdByte(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      byte
		wantError error
	}{
		{"servo id 2", "2", 0x02, nil},
		{"string invalid servo id", "invalid", 0, strconv.ErrSyntax},
		{"out of range servo id", "72", 0, InvalidDriveIdError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := motorIdByte(tt.input)
			if err != nil {
				if errors.Is(err, tt.wantError) != true {
					t.Errorf("want %q; got %q", tt.wantError, err)
				}
			} else {
				if id != tt.want {
					t.Errorf("want %q; got %q", tt.want, id)
				}
				if err != tt.wantError {
					t.Errorf("want %q; got %q", tt.wantError, err)
				}
			}
		})
	}
}

func TestFuncCode(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      byte
		wantError error
	}{
		{"valid stop", "stop", 0x03, nil},
		{"invalid command", "expload", 0, FuncCodeNotImplemented},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fc, err := funcCode(tt.input)
			if err != nil {
				if errors.Is(err, tt.wantError) != true {
					t.Errorf("want %q; got %q", tt.wantError, err)
				}
			} else {
				if fc != tt.want {
					t.Errorf("want %q; got %q", tt.want, fc)
				}
				if err != tt.wantError {
					t.Errorf("want %q; got %q", tt.wantError, err)
				}
			}
		})
	}
}

func TestCreatePacket(t *testing.T) {
	tests := []struct {
		name         string
		inputCommand string
		inputData    int
		want         []byte
		wantError    error
	}{
		{"valid stop command", "stop", 0, []byte{0x02, 0x83}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pkt, err := CreatePacket(tt.inputCommand, tt.inputData)
			if err != nil {
				if errors.Is(err, tt.wantError) != true {
					t.Errorf("want %q; got %q", tt.wantError, err)
				}
			} else {
				if bytes.Equal(pkt, tt.want) != true {
					t.Errorf("want %q; got %q", tt.want, pkt)
				}
				if err != tt.wantError {
					t.Errorf("want %q; got %q", tt.wantError, err)
				}
			}
		})
	}
}
