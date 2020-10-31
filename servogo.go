package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nickherrig/servogo/dmm"

	"github.com/tarm/serial"
)

func main() {

	// command line flags for collecting user input
	id := flag.Int("id", 0, "the motor id, defaults to 0  (Optional)")
	command := flag.String("command", "", "the command to send to the motor (Required)")
	data := flag.Int("data", 0, "data to send with motor command (Optional)")
	flag.Parse()

	// Create dmm packet from command and data
	pkt, err := dmm.CreatePacket(*id, *command, *data)
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
	buf := make([]byte, 128)
	n, err = port.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	// Parse dmm response packet from motor response
	res := buf[:n]
	msg, err := dmm.ParsePacket(res)
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
