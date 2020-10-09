package main

import "testing"

func TestPacketLength(t *testing.T) {
	tests := []struct {
		name      string
		input     int
		want      byte
		wantError error
	}{
		{
			name:      "one positive data packet",
			input:     50,
			want:      0x80,
			wantError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pl, err := packetLength(tt.input)
			if pl != tt.want {
				t.Errorf("want %q; got %q", tt.want, pl)
			}
			if err != tt.wantError {
				t.Errorf("want %q; got %q", tt.wantError, err)
			}
		})
	}
}
