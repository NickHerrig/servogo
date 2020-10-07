package main

import (
	"log"
	//	"os"
	"fmt"
	//	"strconv"

	"go.bug.st/serial"
)

func main() {

	//	p := os.Getenv("SERVO_USB_PORT")
	//	i := os.Getenv("SERVO_DRIVE_ID")
	//
	//	id, err := strconv.Atoi(i)
	//	if err != nil {
	//		log.Fatalf("Failed to convert servo id env var to stringe: %v", err)
	//	}

	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}

	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}

	mode := &serial.Mode{
		BaudRate: 38400,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}

	port, err := serial.Open("/dev/cu.usbserial-DN04SV8H", mode)
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	//forward packet = x02 xe3 xbd xfe xc9 x80 xe9
	//stop packet = x02 x83 x80 x85

	stop := []byte{0x02, 0x83, 0x80, 0x85}

	n, err := port.Write(stop)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(n)
}
