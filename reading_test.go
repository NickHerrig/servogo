package main

import (
	"errors"
	"testing"
)

func TestParseData(t *testing.T) {
	tests := []struct {
		name  string
		data  []byte
		want  int 
	}{
		{"one positive data packet", []byte{0xBC}, 60},
		{"one negative data packet", []byte{0xC4}, -60},
		{"two positive data packet", []byte{0x81, 0x96}, 150},
		{"two negative data packet", []byte{0xFE, 0xEA}, -150},
		{"three positive data packet", []byte{0x80, 0xC6, 0xA8}, 9000},
		{"three negative data packet", []byte{0xFF, 0xB9, 0xD8}, -9000},
		{"four positive data packet", []byte{0xB9, 0x9C, 0x9C, 0x80}, 120000000},
		{"four negative data packet", []byte{0xC6, 0xE3, 0xE4, 0x80}, -120000000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := parseData(tt.data)
			if i != tt.want {
				t.Errorf("want %d; got %d", tt.want, i)
			}
		})
	}
}

func TestValidateChecksum(t *testing.T) {
	tests := []struct {
		name   string
		packet []byte
		want   error
	}{
		{"valid checksum", []byte{0x02, 0x83, 0x01, 0x86}, nil},
		{"invalid checksum", []byte{0x02, 0x83, 0x01, 0x88}, InvalidChecksumError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateChecksum(tt.packet)
			if err != nil {
				if errors.Is(err, tt.want) != true {
					t.Errorf("want %q; got %q", tt.want, err)
				}
			}
		})
	}
}
