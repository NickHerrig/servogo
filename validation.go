package main

import (
    "errors"
)

type Values struct {
    min, max int
}

var MissingCommandError = errors.New("No command sent, Motor command is required")
var InvalidCommandError = errors.New("Invalid command sent, send an valid command")
var InvalidDataError = errors.New("Invalid data for command, Send data in range")

func ValidateInput(c string, d int) error{
    //Map of commands and valid data ranges for input validation
    validInputs := map[string]Values {
        "stop":    {0, 0},
        "forwards": {0, 0},
        "backwards": {0, 0},
    }

    //Check that user passed a command
    if c == "" {
        return MissingCommandError 
    }

    //Check that user passed a valid command
	if  _, ok := validInputs[c]; !ok {
		return InvalidCommandError 
	}

    //Check that user passed valid data for command
    min := validInputs[c].min
    max := validInputs[c].max
    if d < min || d > max {
        return InvalidDataError 
    }
    
    return nil
}
