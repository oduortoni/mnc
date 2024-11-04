package tcp

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"

	"mnc/mnc"
)

func Handle(c *net.Conn, roomMgr *mnc.Rooms) {
	defer (*c).Close()

	var rsuccess bool
	var nbytes int
	buff := make([]byte, 1024)

	reader := bufio.NewReader(*c)
	writer := bufio.NewWriter(*c)
	socket := bufio.NewReadWriter(reader, writer)

	// #begin FETCH_USERNAME
	namePrompt := "Enter a username: "
	write(socket, namePrompt)
	nbytes, rsuccess = read(socket, buff)
	if !rsuccess {
		return
	}
	name := strings.TrimSpace(mnc.ToString(buff[:nbytes]))
	// #end FETCH_USERNAME

	// #begin ROOM_JOINING
	roomsToJoin := "\n--- pick a room to join ----\n" + roomMgr.List() + "\n>_"
	write(socket, roomsToJoin)

	roomID := 0
	nbytes, rsuccess = read(socket, buff)
	if !rsuccess {
		return
	}
	roomID, _ = strconv.Atoi(strings.TrimSpace(mnc.ToString(buff[:nbytes])))
	client := mnc.NewMember(name, roomID, c) // create a client abstraction
	roomID, _ = roomMgr.Join(client, roomID)
	write(socket, fmt.Sprintf("You are joining room number %d\n", roomID))
	currentRoom := roomMgr.Rooms[roomID]
	history := currentRoom.History.List()
	write(socket, history)
	currentRoom.Broadcast(client, fmt.Sprintf("%s joined the room", name), false)
	// #end ROOM_JOINING

	defer func() {
		currentRoom.Leave(client)
		currentRoom.Broadcast(client, fmt.Sprintf("%s has left the room.", name), true)
	}()

	// #begin MESSAGE_LOOP
	for {
		nbytes, rsuccess = read(socket, buff)
		if !rsuccess {
			break // out of the loop to allow for cleanup and prevent another read
		}
		message := strings.TrimSpace(mnc.ToString(buff[:nbytes]))
		fmt.Printf("] MSG [ %q\n", message)

		if message == "" {
			write(socket, "Message cannot be empty.\n")
			continue
		}

		currentRoom.Broadcast(client, message, true)
	}
	// #end MESSAGE_LOOP
}

func read(socket *bufio.ReadWriter, buff []byte) (int, bool) {
	n, err := socket.Read(buff)
	if err != nil {
		errmsg := strings.TrimSpace(err.Error())
		if errmsg != "EOF" {
			fmt.Println("error while reading from client:", errmsg)
		} else {
			fmt.Println("another one leaves")
		}
		return 0, false
	}
	return n, true
}

func write(socket *bufio.ReadWriter, data string) {
	_, err := socket.Write([]byte(data))
	if err != nil {
		errmsg := strings.TrimSpace(err.Error())
		if errmsg != "EOF" {
			fmt.Println("error while reading from client:", errmsg)
		} else {
			fmt.Println("another one leaves")
		}
		return
	}
	socket.Flush()
}
