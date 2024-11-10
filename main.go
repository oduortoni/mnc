package main

import (
	"mnc/mnc"
	"mnc/tcp"
)

const (
	MAXNUMROOMS = 100
	MAXROOMSIZE = 2
)

func main() {
	address := ":9000"
	roomsManager := mnc.NewRooms(MAXNUMROOMS, MAXROOMSIZE)
	tcp.Server(address, roomsManager)
}
