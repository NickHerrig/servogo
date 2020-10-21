package main

import (
	"flag"
	"fmt"
	"log"
	//	"log"
	"os"
	//	"strconv"
	"time"

	"github.com/tarm/serial"
)

func main() {
	// Collect user input with flags
	command := flag.String("command", "", "the command to send to the motor (Required)")
	data := flag.Int("data", 0, "data to send with motor command")
	flag.Parse()

	// Validate user input
	err := ValidateInput(*command, *data)
	if err != nil {
		flag.PrintDefaults()
		log.Fatal(err)
	}
	fmt.Println("Command:", *command, "Data:", *data)

	// Open serial port and defer closing
	port, err := openPort()
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	// Create dmm packet from command and data
	pkt, err := CreatePacket(*command, *data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pkt)

	/*
	       // Write data and register response
	       _, err := port.Write(pkt)
	       if err != nil {
	           log.Fatal(err)
	       }

	       // Read data from drive into buffer
	   	buf := make([]byte, 128)
	   	n, err = s.Read(buf)
	   	if err != nil {
	   		log.Fatal(err)
	   	}

	       // Parse dmm response packet
	       res, err := ParsePacket(n)
	       if err != nil {
	           log.Fatal(err)
	       }

	       // Do something with dmm response
	       TODO
	*/
}

func openPort() (*serial.Port, error) {
	n, ok := os.LookupEnv("SERVO_USB_PORT")
	if !ok {
		log.Fatal("SERVO_USB_PORT env var not set")
	}

	cf := &serial.Config{
		Name:        n,
		Baud:        38400,
		Size:        8,
		ReadTimeout: time.Millisecond * 500,
	}

	p, err := serial.OpenPort(cf)
	if err != nil {
		return nil, err
	}

	return p, nil
}
