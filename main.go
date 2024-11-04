package main

import (
	"fmt"
	"log"
	"net"

	"mnc/mnc"
	"mnc/tcp"
)

const (
	MAXNUMROOMS    = 100
	MAXROOMSIZE = 2 // return to 10
)

func main() {
	laddress := ":9000"
	roomsManager := mnc.NewRooms(MAXNUMROOMS, MAXROOMSIZE)

	l, e := net.Listen("tcp", laddress)
	if e != nil {
		log.Fatalln("Error creating server")
	}
	fmt.Printf("\n\n\tServer listening on %s\n\n\n", laddress)

	for {
		c, e := l.Accept()
		if e != nil {
			log.Fatalln("Error accepting client connection")
		}
		go tcp.Handle(&c, roomsManager)
	}
}
