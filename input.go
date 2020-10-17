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

    if c == "" {
        return errors.New("Servo command is required!")
    }

	if  _, ok := validInputs[c]; !ok {
		return errors.New("Command not Implemeneted") 
	}
    
    return nil
}
