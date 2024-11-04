package mnc

import (
	"fmt"
	"sync"
	"time"
)

type Room struct {
	Id       int
	Name     string
	Capacity int
	Members  map[string]*Member
	History  *History
}

func NewRoom(id int, name string, capacity int) *Room {
	return &Room{
		Id:       id,
		Name:     name,
		Capacity: capacity,
		Members:  make(map[string]*Member),
		History:  &History{},
	}
}

func (r *Room) Join(member *Member) (bool, int) {
	if len(r.Members) >= r.Capacity {
		return false, FULL // Room is full
	}

	// Check if the member is already in the room
	if _, exists := r.Members[member.Name]; exists {
		fmt.Printf("User %s already exists\n", member.Name)
		return false, EXISTS
	}

	// Add the member to the room
	r.Members[member.Name] = member
	member.RoomID = r.Id // Store the room ID for the member
	return true, SUCCESS
}

func (r *Room) Leave(member *Member) bool {
	if _, exists := r.Members[member.Name]; exists {
		delete(r.Members, member.Name)
		return true
	}
	return false
}

var mutex sync.Mutex

func (r *Room) Broadcast(sender *Member, message string, save bool) {
	formattedMsg := formatMessage(sender.Name, message)

	if save {
		r.History.Save(formattedMsg)
	}

	mutex.Lock()
	defer mutex.Unlock()

	for _, member := range r.Members {
		if member != sender {
			_, err := (*member.Entity).Write([]byte(formattedMsg))
			if err != nil {
				fmt.Printf("Error sending message to %s: %v\n", member.Name, err)
			}
		}
	}
}

func formatMessage(sendersName, message string) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	if sendersName == "" {
		return fmt.Sprintf("[%s] %s\n", timestamp, message)
	}

	// Regular chat message
	return fmt.Sprintf("[%s][%s]: %s\n", timestamp, sendersName, message)
}
