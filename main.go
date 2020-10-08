package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/tarm/serial"
)

func main() {

	data := flag.Int("data", 1, "numerical data to pass to the servo drive")
	flag.Parse()

	p := os.Getenv("SERVO_USB_PORT")
	i := os.Getenv("SERVO_DRIVE_ID")
	id, err := strconv.Atoi(i)
	if err != nil {
		log.Fatalf("Failed to convert servo id env var to stringe: %v", err)
	}
	bid := byte(id)

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
		log.Fatal(err)
	}

	buf := make([]byte, 128)
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("%q", buf[:n])

}
