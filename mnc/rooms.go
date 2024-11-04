package mnc

import (
	"fmt"
	"strings"
)


type Rooms struct {
	CurrentNumber int
	MaxNumRooms int
    MaxRoomSize int
	Rooms         []*Room
}

func NewRooms(maxnumrooms, maxroomsize int) *Rooms {
    rooms :=  &Rooms{
        CurrentNumber: 0,
        MaxNumRooms: maxnumrooms,
        MaxRoomSize: maxroomsize,
        Rooms:         []*Room{},
    }
    main := NewRoom(0, "main", maxroomsize)
    rooms.CurrentNumber++
    rooms.Rooms = append(rooms.Rooms, main)
    return rooms
}


func (rooms *Rooms) CreateRoom(name string, capacity int) (int, error) {
    if rooms.CurrentNumber >= rooms.MaxNumRooms {
        return -1, fmt.Errorf("maximum number of rooms reached")
    }

    name = strings.TrimSpace(name)
    if len(name) == 0 {
        name = fmt.Sprintf("Room#%d", rooms.CurrentNumber)
    }

    uniqueName := name
    for _, room := range rooms.Rooms {
        if room.Name == uniqueName {
            uniqueName = fmt.Sprintf("%s_%d", name, rooms.CurrentNumber)
        }
    }

    room := NewRoom(rooms.CurrentNumber, uniqueName, rooms.MaxRoomSize)
    rooms.Rooms = append(rooms.Rooms, room)
    rooms.CurrentNumber++ // next room's ID

    return room.Id, nil
}

func (rooms *Rooms) Join(member *Member, roomId int) (int, int) {
    if roomId < 0 || roomId >= rooms.CurrentNumber {
        return -1, IDRANGE // invalid room ID
    }

    room := rooms.Rooms[roomId]
    success, status := room.Join(member)
    if success {
        return room.Id, SUCCESS
    }

    // member already exists in the room
    if status == EXISTS {
        return -1, EXISTS
    }

    // try to join another available room
    for _, r := range rooms.Rooms {
        if len(r.Members) < r.Capacity {
            if success, _ := r.Join(member); success {
                return r.Id, SUCCESS
            }
        }
    }

    // no available rooms, create a new one
    newRoomId, err := rooms.CreateRoom("", 5)
    if err != nil {
        return -1, FULL // could not create a new room
    }

    // join the member to the new room
    newRoom := rooms.Rooms[newRoomId]
    if success, _ := newRoom.Join(member); success {
        return newRoom.Id, SUCCESS // return the new room ID with success
    }
    return -1, FULL // in case of failure to join the new room
}


func (r *Rooms)List() string {
    roomList := ""
    for index, room := range r.Rooms {
        roomList += fmt.Sprintf("%d) %s\n", index, room.Name)
    }
    return roomList
}
