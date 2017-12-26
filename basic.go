package main

import (
	"fmt"
	"github.com/tarm/serial"
	"log"
	"time"
)

func main() {
	// c := &serial.Config{Name: "COM4", Baud: 115200}
	c := &serial.Config{
		Name:        "COM4",
		Baud:        128000, //115200,
		Size:        8,
		StopBits:    serial.Stop1,
		Parity:      serial.ParityNone,
		ReadTimeout: 500 * time.Millisecond,
	}

	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	// n, err := s.Write([]byte("test"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	buf := make([]byte, 128)
	for {
		n, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("N is :%v\n", n)
		fmt.Printf("%v", string(buf[:n]))
		// fmt.Printf("%v", string(buf))
		// fmt.Printf("%v", buf)
	}
}
