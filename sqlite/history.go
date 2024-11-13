package sqlite

import (
	"database/sql"
	"log"

	"mnc/mnc"

	_ "github.com/mattn/go-sqlite3"
)

func CreateHistoryTable(db *sql.DB, args ...any) any {
	query, ok := args[0].(string)
	if ok {
		_, err := db.Exec(query)
		return err == nil
	}
	return false
}

/*
* create a room
* args: query, content roomname
 */
func HistoryInsert(db *sql.DB, args ...any) any {
	query, ok := args[0].(string)
	if ok {
		content, ok := args[1].(string)
		if ok {
			roomname, ok := args[2].(string)
			if ok {
				_, err := db.Exec(query, content, roomname)
				return err == nil
			}
		}
	}
	return false
}

/*
* select all rooms from the database
 */
func HistorySelectAll(db *sql.DB, args ...any) any {
	query, ok := args[0].(string)
	if !ok {
		return nil
	}
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	history := &mnc.History{}
	// iterate over the rows appending a room to rooms
	for rows.Next() {
		var id int
		var content string
		var roomname string
		if err := rows.Scan(&id, &content, &roomname); err != nil {
			log.Fatal(err)
		}
		message := &mnc.Message{
			RoomName: roomname,
			Content:  content,
		}
		history.Messages = append(history.Messages, message)
	}
	if err := rows.Err(); err != nil {
		return nil
	}
	return history
}

/*
* select a room from the database by id
 */
func HistorySelectByRoomName(db *sql.DB, args ...any) any {
	query, ok := args[0].(string)
	if !ok {
		return nil
	}
	roomNameArg, ok := args[1].(string)
	if !ok {
		return nil
	}
	rows, err := db.Query(query, roomNameArg)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	history := &mnc.History{}
	// iterate over the rows appending a room to rooms
	for rows.Next() {
		var id int
		var content string
		var roomname string
		if err := rows.Scan(&id, &content, &roomname); err != nil {
			log.Fatal(err)
		}
		message := &mnc.Message{
			RoomName: roomname,
			Content:  content,
		}
		history.Messages = append(history.Messages, message)
	}
	if err := rows.Err(); err != nil {
		return nil
	}
	return history
}

/*
* deletes an piece of history
 */
func HistoryDelete(db *sql.DB, args ...any) any {
	query, ok := args[0].(string)
	if ok {
		historyId, ok := args[1].(int)
		if ok {
			_, err := db.Exec(query, historyId)
			return err == nil
		}
	}
	return false
}
