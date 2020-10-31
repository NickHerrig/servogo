package dmm

import (
	"errors"
)

// Errors returned from CreatePacket() input validation
var (
	MissingCommandError = errors.New("No command sent, motor command is required")
	InvalidCommandError = errors.New("Invalid command sent, send an valid command")
	InvalidDataError    = errors.New("Invalid data sent for command")
	InvalidServoIdError = errors.New("Invalid servo id sent, must be 0~64")
)

func validateInput(servoId int, command string, data int) error {

	// Validate servoId is between 0 and 63!
	if servoId < 0 || servoId > 63 {
		return InvalidDriveIdError
	}

	if command == "" {
		return MissingCommandError
	}

	//Check that user passed a command
	if command == "" {
		return MissingCommandError
	}

	//Check that user passed a valid command
	if _, ok := commandMap[command]; !ok {
		return InvalidCommandError
	}

	//Check that user passed valid data for command
	min := commandMap[command].min
	max := commandMap[command].max
	if data < min || data > max {
		return InvalidDataError
	}

	return nil
}
