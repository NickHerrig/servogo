package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/tarm/serial"
)

func main() {
	// Collect user input
	command := flag.String("command", "", "the command to send to the motor (Required)")
	data := flag.Int("data", 0, "data to send with motor command (Optional)")
	flag.Parse()

	// Validate user input, if error, print flag details and log failure
	err := ValidateInput(*command, *data)
	if err != nil {
		flag.PrintDefaults()
		log.Fatal(err)
	}

	// If command was passed with no data, change data to record in map[command]Functions.data
	// example: "servogo forwards" == servogo fowards --data 13000000
	if *data == 0 {
		*data = commandMap[*command].data
	}

	// Create dmm packet from command and data
	pkt, err := CreatePacket(*command, *data)
	if err != nil {
		log.Fatal(err)
	}

	//Open serial port and defer closing
	port, err := openPort()
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	// Write data to serial port
	n, err := port.Write(pkt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Bytes Written:", n)

	// Read data from drive into buffer
	buf := make([]byte, 8)
	n, err = port.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	// Parse dmm response packet from motor response
	res := buf[:n]
	msg, err := ParsePacket(res)
	if err != nil {
		log.Fatal(err)
	}

	// Do something with dmm response
	fmt.Println("Running command:", *command)
	fmt.Println(msg)
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
