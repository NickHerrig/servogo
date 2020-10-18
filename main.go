package main

import (
	"fmt"
    "flag"
    "log"
//	"log"
//	"os"
//	"strconv"
//	"time"

//	"github.com/tarm/serial"
)

func main() {

    command := flag.String("command", "", "the command to send to the motor (Required)")
    data := flag.Int("data", 0, "data to send with motor command")
    flag.Parse()

    err := ValidateInput(*command, *data)
    if err != nil {
        flag.PrintDefaults()
        log.Fatal(err)
    }

    fmt.Println("Command:", *command, "Data:", *data)


/*  SKETCHING DATA FLOW AND LAYOUT
    INPUT:
        command and data _should_ be  English... Implementation will differ by motor
        keep flags in main...

    INPUT VALIDATION:
        Input validation will _probably_ be motor specific.
        maybe this goes in a module?

    SERIAL CONFIG:
        Serial configuration is the only external dependency.
        Most packages come with a Config struct for BaudRates/Ports/etc..
        Keep in main for now

    CREATE PACKET:
        Once we have valid input, we need to craft a packet...
        packets are motor specific, but I want to keep I/O in main..
        is that a valid want? for now I think so.
        Does this belong in a module specific to motors? along with input validation
        and packet parsing? Might be interesting to look into a motor struct and implement
        an interface? similar to past projects where storage is an interface for mysql/sqlite3

    WRITE PACKET:
        Pretty straight forward... once you have the packet, write serial to port..

    READ PACKET:
        Also straightforward, need to parse to make sense of it!

    PARSE PACKET:
        Parsing the packet will also differ deppending on the motor...
        We will read packet here, but need to pass the packet to module to make sense of it..
        more I think of it, the motor specifics like creating packet, validating data(input),
        and reading responses probably belong in a module.. can I implement an interface
        to do this elegantly if I have more than one motor? is this a pre-optimization?



	cf := &serial.Config{
		Name:        p,
		Baud:        38400,
		Size:        8,
		ReadTimeout: time.Millisecond * 500,
	}

	s, err := serial.OpenPort(cf)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

*/
}
