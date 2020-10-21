package main

type Functions struct {
	min      int
	max      int
	function string
	code     byte
	data     int
}

//Map of commands and valid data ranges for input validation
var commandMap = map[string]Functions{
	"stop":      {0, 0, "Go_Relative_Pos", 0x03, 0},
	"forwards":  {0, 0, "Go_Relative_Pos", 0x03, 13000000},
	"backwards": {0, 0, "Go_Relative_Pos", 0x03, -13000000},
}
