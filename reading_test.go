package main

import (
	"errors"
	"testing"
)

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
