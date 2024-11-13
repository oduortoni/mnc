package sqlite

import (
	"database/sql"
	"log"

	"mnc/mnc"

	_ "github.com/mattn/go-sqlite3"
)

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
				roomDescription, ok := args[3].(string)
				if ok {
					_, err := db.Exec(query, roomName, roomCapacity, roomDescription)
					return err == nil
				}
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
		var description string
		if err := rows.Scan(&id, &name, &capacity, &description); err != nil {
			log.Fatal(err)
		}
		room := &mnc.Room{
			Id:          id,
			Name:        name,
			Description: description,
			Capacity:    capacity,
		}
		rooms = append(rooms, room)
	}
	if err := rows.Err(); err != nil {
		return nil
	}
	return rooms
}

/*
* select a room from the database by id
 */
func RoomSelectById(db *sql.DB, args ...any) any {
	query, ok := args[0].(string)
	if !ok {
		return nil
	}
	roomId, ok := args[1].(int)
	if !ok {
		return nil
	}
	rows, err := db.Query(query, roomId)
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
		var description string
		if err := rows.Scan(&id, &name, &capacity, &description); err != nil {
			log.Fatal(err)
		}
		room := &mnc.Room{
			Id:          id,
			Name:        name,
			Description: description,
			Capacity:    capacity,
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
				roomDescription, ok := args[3].(string)
				if ok {
					_, err := db.Exec(query, roomId, roomName, roomDescription)
					return err == nil
				}
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
