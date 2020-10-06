package main

import (
	"log"
	"os"
	"strconv"

	"github.com/jacobsa/go-serial/serial"
)

func main() {

	p := os.Getenv("SERVO_USB_PORT")
	i := os.Getenv("SERVO_DRIVE_ID")

	id, err := strconv.Atoi(i)
	if err != nil {
		log.Fatalf("Failed to convert servo id env var to stringe: %v", err)
	}

	options := serial.OpenOptions{
		PortName:        p,
		BaudRate:        38400,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	defer port.Close()

	pkt = []byte{}

	n, err := port.Write(pk)
	if err != nil {
		log.Fatalf("port.Write: %v", err)
	}

}
