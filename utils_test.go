package main

import (
	"testing"
)

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
