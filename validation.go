package main

import (
	"errors"
)

var MissingCommandError = errors.New("No command sent, motor command is required")
var InvalidCommandError = errors.New("Invalid command sent, send an valid command")
var InvalidDataError = errors.New("Invalid data sent for command")

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
