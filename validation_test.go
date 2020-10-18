package main

import (
    "testing"
    "errors"
)

func TestInputValidation(t *testing.T) {
	tests := []struct {
		name      string
		command   string 
		data      int 
		want      error
	}{
		{"Valid stop", "stop", 0, nil},
		{"Invalid stop", "stop", 10, InvalidDataError},
		{"No command", "", 100, MissingCommandError},
		{"Invalid command", "expload", 0, InvalidCommandError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateInput(tt.command, tt.data)
			if err != nil {
                if errors.Is(err, tt.want) != true {
                    t.Errorf("want %q; got %q", tt.want, err)
                }
			}
		})
	}
}
