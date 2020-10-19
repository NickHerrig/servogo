package main

import (
	"errors"
)

type Functions struct {
	min      int
	max      int
	function string
	code     byte
}

var MissingCommandError = errors.New("No command sent, motor command is required")
var InvalidCommandError = errors.New("Invalid command sent, send an valid command")
var InvalidDataError = errors.New("Invalid data sent for command")

//Map of commands and valid data ranges for input validation
var commandMap = map[string]Functions{
	"stop":      {0, 0, "Go_Relative_Pos", 0x03},
	"forwards":  {0, 0, "Go_Relative_Pos", 0x03},
	"backwards": {0, 0, "Go_Relative_Pos", 0x03},
}

func ValidateInput(c string, d int) error {
	//Check that user passed a command
	if c == "" {
		return MissingCommandError
	}

	//Check that user passed a valid command
	if _, ok := commandMap[c]; !ok {
		return InvalidCommandError
	}

	//Check that user passed valid data for command
	min := commandMap[c].min
	max := commandMap[c].max
	if d < min || d > max {
		return InvalidDataError
	}

	return nil
}
