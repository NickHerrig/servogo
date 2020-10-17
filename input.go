package main

import (
    "errors"
)

func ValidateInput(c string, d int) error{
    if c == "" {
        return errors.New("Servo command is required!")
    }
    return nil
}
