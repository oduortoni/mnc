package main

import (
	"fmt"
	"mnc/http"
	"mnc/mnc"
)

const (
	MAXNUMROOMS = 100
	MAXROOMSIZE = 2
)

func main() {
	address := ":9000"
	fmt.Printf("\n\n\tServer running on %s\n\n", address)
	roomsManager := mnc.NewRooms(MAXNUMROOMS, MAXROOMSIZE)
	http.Server(address, roomsManager)
}

// CREATE TABLES START

// package main

// import "mnc/sqlite"

// func main() {
// 	sqlite.Run(sqlite.CreateRoomsTable, sqlite.RoomsCreateTableQuery)
// 	roomName := "main"
// 	description := "Main room is a room that is dedicated to all communications between the establishment and the outside world. Any official communication will be through the main room"
// 	maxMembers := 5
// 	sqlite.Run(sqlite.RoomCreate, sqlite.RoomsInsertQuery, roomName, maxMembers, description)
// 	sqlite.Run(sqlite.CreateHistoryTable, sqlite.HistoryCreateTableQuery)
// 	content := "Today is a lovely day to begin MNCiers!"
// 	sqlite.Run(sqlite.HistoryInsert, sqlite.HistoryInsertQuery, content, roomName)
// }

// CREATE TABLES END

// package main

// import (
// 	"mnc/sqlite"
// )

// func main() {
// 	sqlite.Run(sqlite.CreateRoomsTable, sqlite.RoomsCreateTableQuery)

// insertSQL := `INSERT INTO rooms(name, capacity) VALUES (?, ?)`
// sqlite.Run(sqlite.RoomCreate, insertSQL, "main", 1)
// sqlite.Run(sqlite.RoomCreate, insertSQL, "vengeance", 90)
// sqlite.Run(sqlite.RoomCreate, insertSQL, "idlers", 21)
// sqlite.Run(sqlite.RoomCreate, insertSQL, "mboggi", 7)

// selectSQL := "SELECT id, name, capacity FROM rooms"
// rooms, ok := sqlite.Run(sqlite.RoomSelect, selectSQL).([]*mnc.Room)
// if ok {
// 	fmt.Println("Rooms: ", rooms)
// }

// selectByIdSQL := "SELECT id, name, capacity, description FROM rooms WHERE id = ?"
// rooms, ok := sqlite.Run(sqlite.RoomSelectById, selectByIdSQL, 3).([]*mnc.Room)
// if ok {
// 	fmt.Println("Room: ", rooms[0].Name)
// }

// updateSQL := `UPDATE rooms SET capacity = ? WHERE name = ?`
// _, updated := sqlite.Run(sqlite.RoomUpdate, updateSQL, 3002, "mboggi").(bool)
// if updated {
// 	fmt.Println("SUCCESS update")
// }

// deleteSQL := `DELETE FROM rooms WHERE name = ?`
// _, deleted := sqlite.Run(sqlite.RoomDelete, deleteSQL, "main").(bool)
// if deleted {
// 	fmt.Println("DELETE success")
// }
// }
