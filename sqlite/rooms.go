package sqlite

import (
	"database/sql"
	"fmt"
	"log"

	"mnc/mnc"

	_ "github.com/mattn/go-sqlite3"
)

/*
* used to open a database connection and invoking the callback that does the real work
 */
func Run(callback func(db *sql.DB, args ...any) any, args ...any) any {
	db, err := sql.Open("sqlite3", "./database/mnc.db") // open or create if it doesnt exist
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return callback(db, args...)
}

func CreateRoomsTable(db *sql.DB, args ...any) any {
	query, ok := args[0].(string)
	if ok {
		_, err := db.Exec(query)
		return err == nil
	}
	return false
}

/*
* create a room
* args: query, name, capacity
 */
func RoomCreate(db *sql.DB, args ...any) any {
	query, ok := args[0].(string)
	if ok {
		roomName, ok := args[1].(string)
		if ok {
			roomCapacity, ok := args[2].(int)
			if ok {
				_, err := db.Exec(query, roomName, roomCapacity)
				return err == nil
			}
		}
	}
	return false
}

/*
* select all rooms from the database
 */
func RoomSelect(db *sql.DB, args ...any) any {
	query, ok := args[0].(string)
	if !ok {
		return nil
	}
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	rooms := []*mnc.Room{}
	// iterate over the rows appending a room to rooms
	for rows.Next() {
		var id int
		var name string
		var capacity int
		if err := rows.Scan(&id, &name, &capacity); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Capacity: %d\n", id, name, capacity)
		room := &mnc.Room{
			Id:       id,
			Name:     name,
			Capacity: capacity,
		}
		rooms = append(rooms, room)
	}
	if err := rows.Err(); err != nil {
		return nil
	}
	return rooms
}

/*
* updates the details of a room
 */
func RoomUpdate(db *sql.DB, args ...any) any {
	query, ok := args[0].(string)
	if ok {
		roomId, ok := args[1].(int)
		if ok {
			roomName, ok := args[2].(string)
			if ok {
				_, err := db.Exec(query, roomId, roomName)
				return err == nil
			}
		}
	}
	return false
}

/*
* deletes an entire room
 */
func RoomDelete(db *sql.DB, args ...any) any {
	query, ok := args[0].(string)
	if ok {
		roomName, ok := args[1].(string)
		if ok {
			_, err := db.Exec(query, roomName)
			return err == nil
		}
	}
	return false
}
