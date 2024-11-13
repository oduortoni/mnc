package sqlite

const  (
	RoomsCreateTableQuery = "CREATE TABLE IF NOT EXISTS rooms (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, capacity INTEGER);"
	RoomsSelectAllQuery = "SELECT id, name, capacity FROM rooms"
	RoomsSelectByIdQuery = "SELECT id, name, capacity FROM rooms WHERE id = ?"
	RoomsInsertQuery = `INSERT INTO rooms(name, capacity) VALUES (?, ?)`
	RoomsUpdateQuery = `UPDATE rooms SET capacity = ? WHERE name = ?`
	RoomsDeleteQuery = `DELETE FROM rooms WHERE name = ?`
)
