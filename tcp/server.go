package tcp

import (
	"fmt"
	"log"
	"net"
	"mnc/mnc"
)

func Server(address string, roomsManager *mnc.Rooms) {
	l, e := net.Listen("tcp", address)
	if e != nil {
		log.Fatalln("Error creating server")
	}
	fmt.Printf("\n\n\tServer listening on %s\n\n\n", address)

	for {
		c, e := l.Accept()
		if e != nil {
			log.Fatalln("Error accepting client connection")
		}
		go Handle(&c, roomsManager)
	}
}
