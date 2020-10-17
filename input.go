package main

import (
    "errors"
)

type Values struct {
    min, max int
}

func ValidateInput(c string, d int) error{

    validInputs := map[string]Values {
        "stop": {0, 0},
    }

    //Check that user passed a command
    if c == "" {
        return errors.New("Servo command is required!")
    }

    //Check that user passed a valid command
	if  _, ok := validInputs[c]; !ok {
		return errors.New("Command not Implemeneted") 
	}

    //Check that user passed valid data for command
    //TODO
    
    return nil
}
