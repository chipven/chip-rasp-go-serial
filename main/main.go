package main

import (
	"fmt"
	"github.com/tarm/goserial"
	"log"
)

func main() {
	c := &serial.Config{Name: "/dev/ttyS0", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	_, err = s.Write([]byte("test"))
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 128)
	fmt.Println(len(buf))

	n, err := s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("%q", buf[:n])
}
