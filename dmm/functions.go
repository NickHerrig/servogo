package dmm

// Functions struct groups dmm funcs and implementation details
type Functions struct {
	function string // the dmm motor function
	min      int    // the minimum data value for a command
	max      int    // the maximum data value for a command
	code     byte   // the dmm motor function code in hexadecimal
	data     int    // the data for a particular command if 0
}

//Map of commands dmm function implementation detail struct
var commandMap = map[string]Functions{
	"stop":       {"Go_Relative_Pos", 0, 0, 0x03, 0},
	"forwards":   {"Go_Relative_Pos", 0, 0, 0x03, 13000000},
	"backwards":  {"Go_Relative_Pos", 0, 0, 0x03, -13000000},
	"send-to":    {"Go_Absolute_Pos", -134217728, 1342177270, 0x01, 0},
	"position":   {"General_Read", 0, 0, 0x0E, 0x1B},
	"set-speed":  {"Set_HighSpeed", 0, 127, 0x14, 0},
	"read-speed": {"Read_HighSpeed", 0, 0, 0x1C, 0},
	"status":     {"RegisterRead_Drive_Status", 0, 0, 0x09, 0},
}
