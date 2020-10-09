package main

import (
	"bytes"
	"errors"
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
		{"two positive data packets", 8000, 0xa0, nil},
		{"two negative data packets", -8000, 0xa0, nil},
		{"three positive data packets", 1000000, 0xc0, nil},
		{"three negative data packets", -1000000, 0xc0, nil},
		{"four positive data packets", 130000000, 0xe0, nil},
		{"four negative data packets", -130000000, 0xe0, nil},
		{"five positive data packets", 140000000, 0, errors.New("data out of range for dmm servo!")},
		{"five negative data packets", -140000000, 0, errors.New("data out of range for dmm servo!")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pl, err := packetLength(tt.input)

			if err != nil {
				if errors.Is(err, tt.wantError) {
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
		{"one positive data packet", 60, []byte{0xbc}, nil},
		{"one negative data packet", -60, []byte{0xc4}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := dataBytes(tt.input)

			if err != nil {
				if errors.Is(err, tt.wantError) {
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
