package dmm 

import (
	"errors"
	"testing"
)

func TestInputValidation(t *testing.T) {
	tests := []struct {
		name    string
		id      int 
		command string
		data    int
		want    error
	}{
		{"Valid stop", 0, "stop", 0, nil},
		{"Invalid stop", 0, "stop", 10, InvalidDataError},
		{"No command", 0, "", 100, MissingCommandError},
		{"Invalid command", 0, "expload", 0, InvalidCommandError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateInput(tt.id, tt.command, tt.data)
			if err != nil {
				if errors.Is(err, tt.want) != true {
					t.Errorf("want %q; got %q", tt.want, err)
				}
			}
		})
	}
}
