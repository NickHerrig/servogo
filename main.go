package main

import (
	"log"
	"os"

	"github.com/jacobsa/go-serial/serial"
)

func main() {

	p := os.Getenv("PYSERVO_USB_PORT")

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

}
