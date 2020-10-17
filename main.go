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

/*
	f, err := funcCode(c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f)

	i, ok := os.LookupEnv("SERVO_DRIVE_ID")
	if !ok {
		log.Fatal("SERVO_DRIVE_ID env var not set")
	}
	id, err := strconv.Atoi(i)
	if err != nil {
		log.Fatalf("Failed to convert servo id env var to string: %v", err)
	}
	idb := byte(id)

	p, ok := os.LookupEnv("SERVO_USB_PORT")
	if !ok {
		log.Fatal("SERVO_USB_PORT env var not set")
	}

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

	packet := []byte{idb}

	n, err := s.Write(packet)
	if err != nil {
		log.Fatalf("Error Writing: %v", err)
	}

	buf := make([]byte, 128)
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(buf[:n])
*/
}
