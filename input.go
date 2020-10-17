package main

import (
    "errors"
)

type Values struct {
    min, max int
}

func ValidateInput(c string, d int) error{

    validInputs := map[string]Values {
        "stop":    {0, 0},
        "forwards": {0, 0},
        "backwards": {0, 0},
    }

    //Check that user passed a command
    if c == "" {
        return errors.New("Motor command is required")
    }

    //Check that user passed a valid command
	if  _, ok := validInputs[c]; !ok {
		return errors.New("Command not Implemeneted") 
	}

    //Check that user passed valid data for command
    min := validInputs[c].min
    max := validInputs[c].max
    if d < min || d > max {
        return errors.New("Invalid data for command")
    }
    
    return nil
}
