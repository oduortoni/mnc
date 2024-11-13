package rooms

import (
	"mnc/sqlite"
)

func Create(name string, capacity int) {
	sqlite.Run(sqlite.RoomCreate, sqlite.RoomsInsertQuery, name, capacity)
}
