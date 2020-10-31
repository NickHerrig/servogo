package dmm 

import (
	"errors"
)

// Errors returned from CreatePacket() input validation
var (
	MissingCommandError = errors.New("No command sent, motor command is required")
	InvalidCommandError = errors.New("Invalid command sent, send an valid command")
	InvalidDataError    = errors.New("Invalid data sent for command")
	InvalidServoIdError    = errors.New("Invalid servo id sent, must be 0~64")
)

func ValidateInput(servoId int, command string, data int) error {

	// Validate servoId is between 0 and 63!
	if id < 0 || id > 63 {
		return 0, InvalidDriveIdError
	}

	if command == "" {
		return MissingCommandError
	}

	//Check that user passed a command
	if command == "" {
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
