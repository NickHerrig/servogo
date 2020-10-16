package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/tarm/serial"
)

func main() {
	if len(os.Args) == 2 {
		c := os.Args[1]
		d := byte(0)
		fmt.Println("Command:", c, "Data:", d)
	} else if len(os.Args) == 3 {
		c := os.Args[1]
		ds := os.Args[2]
		d, err := strconv.Atoi(ds)
		if err != nil {
			log.Fatalf("Positional arg 2 must be an integer.")
		}
		fmt.Println("Command: ", c, "Data: ", d)
	} else {
		log.Fatal("Error: must folow format {command} or {command} {data}, example 'servo send-to 2000'")
	}

	i := os.Getenv("SERVO_DRIVE_ID")
	id, err := strconv.Atoi(i)
	if err != nil {
		log.Fatalf("Failed to convert servo id env var to string: %v", err)
	}
	bid := byte(id)

	p := os.Getenv("SERVO_USB_PORT")
	c := &serial.Config{
		Name:        p,
		Baud:        38400,
		Size:        8,
		ReadTimeout: time.Millisecond * 500,
	}

	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	//forward packet = x02 xe3 xbd xfe xc9 x80 xe9
	//stop packet = x02 x83 x80 x85

	stop := []byte{bid, 0x83, 0x80, 0x85}

	n, err := s.Write(stop)
	if err != nil {
		log.Fatalf("Error Writing: %v", err)
	}

	buf := make([]byte, 128)
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(buf[:n])
}
