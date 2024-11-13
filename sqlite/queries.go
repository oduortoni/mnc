package sqlite

const  (
	RoomsCreateTableQuery = "CREATE TABLE IF NOT EXISTS rooms (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, capacity INTEGER, description TEXT);"
	RoomsSelectAllQuery = "SELECT id, name, capacity, description FROM rooms"
	RoomsSelectByIdQuery = "SELECT id, name, capacity, description FROM rooms WHERE id = ?"
	RoomsInsertQuery = `INSERT INTO rooms(name, capacity, description) VALUES (?, ?, ?)`
	RoomsUpdateQuery = `UPDATE rooms SET capacity = ? AND description = ? WHERE name = ?`
	RoomsDeleteQuery = `DELETE FROM rooms WHERE name = ?`
	HistoryCreateTableQuery = "CREATE TABLE IF NOT EXISTS history (id INTEGER PRIMARY KEY AUTOINCREMENT, content TEXT, room TEXT);"
	HistorySelectByRoomNameQuery = "SELECT id, content, room FROM history WHERE room = ?"
	HistoryInsertQuery = `INSERT INTO history(content, room) VALUES (?, ?)`
	HistoryDeleteQuery = `DELETE FROM history WHERE id = ?`
)
