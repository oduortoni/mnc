package http

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"mnc/mnc"
	"mnc/sqlite"
)

type ReceivedMessage struct {
	RoomName string `json:"roomname"`
	Message  string `json:"message"`
}

type ResponseMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func SaveMessage(roomsManager *mnc.Rooms) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			// TODO: Handle errors gracefully
			message := ResponseMessage{
				Status:  "fail",
				Message: "Method not allowed!",
			}
			messageJson, _ := json.Marshal(message)
			w.Write(messageJson)
			return
		}

		var req ReceivedMessage
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		if err != nil {
			// TODO: Handle errors gracefully
			message := ResponseMessage{
				Status:  "fail",
				Message: "Failed to decode message!",
			}
			messageJson, _ := json.Marshal(message)
			w.Write(messageJson)
			return
		}

		// Log or process the data as needed
		fmt.Printf("Received message for room '%s': %s\n", req.RoomName, req.Message)

		// save to database
		sqlite.Run(sqlite.HistoryInsert, sqlite.HistoryInsertQuery, req.Message, req.RoomName)

		// save to the room history
		room := roomsManager.GetByName(req.RoomName)
		if room != nil {
			message := &mnc.Message{
				RoomName: req.RoomName,
				Content:  req.Message,
			}
			room.History.Save(message)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// TODO: Handle errors gracefully
		message := ResponseMessage{
			Status:  "success",
			Message: "Message saved successfully",
		}
		messageJson, _ := json.Marshal(message)
		w.Write(messageJson)
	}
}

func CreateRoom(roomsManager *mnc.Rooms) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roomName := r.FormValue("name")
		roomCapacityStr := r.FormValue("capacity")
		roomCapacity, err := strconv.Atoi(roomCapacityStr)
		if err != nil {
			roomCapacity = 10
		}
		roomDescription := r.FormValue("description")

		sqlite.Run(sqlite.RoomCreate, sqlite.RoomsInsertQuery, roomName, roomCapacity, roomDescription)

		roomsManager.CreateRoom(roomName, roomCapacity, roomDescription)

		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func ExploreOne(roomsManager *mnc.Rooms) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		split := strings.Split(r.URL.Path, "/")
		nm1 := len(split) - 1
		var id int
		id, err := strconv.Atoi(split[nm1])
		if err != nil {
			id = 1
		}

		dbrooms := sqlite.Run(sqlite.RoomSelectById, sqlite.RoomsSelectByIdQuery, id).([]*mnc.Room)
		if dbrooms == nil {
			errRooms := mnc.Room{}
			errJson, _ := json.Marshal(errRooms)
			w.Write(errJson)
			return
		}
		if len(dbrooms) == 0 {
			errRooms := mnc.Room{}
			errJson, _ := json.Marshal(errRooms)
			w.Write(errJson)
			return
		}
		roomsManager.Rooms = dbrooms
		dbRoom := dbrooms[0]

		// fetch all the history of the room
		history := sqlite.Run(sqlite.HistorySelectByRoomName, sqlite.HistorySelectByRoomNameQuery, dbRoom.Name).(*mnc.History)
		if history == nil {
			errHistory := mnc.History{}
			errJson, _ := json.Marshal(errHistory)
			w.Write(errJson)
			return
		}
		if len(history.Messages) == 0 {
			errHistory := mnc.Room{}
			errJson, _ := json.Marshal(errHistory)
			w.Write(errJson)
			return
		}
		dbRoom.History = history

		// IF WE WERE TO CREATE AN API
		// w.Header().Set("Content-Type", "application/json")

		// fmt.Printf("ROOM: %v\n", dbRoom)
		// fmt.Printf("ROOM: %v\n", dbRoom.Name)

		// roomsJson, err := json.Marshal(dbRoom)
		// if err != nil {
		// 	errRooms := mnc.Room{}
		// 	errJson, _ := json.Marshal(errRooms)
		// 	w.Write(errJson)
		// } else {
		// 	w.Write(roomsJson)
		// }

		t := template.Must(template.ParseFiles("templates/room.html"))
		t.Execute(w, dbRoom)
	}
}

func ExploreAll(roomsManager *mnc.Rooms) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dbrooms, ok := sqlite.Run(sqlite.RoomSelect, sqlite.RoomsSelectAllQuery).([]*mnc.Room)
		if ok {
			roomsManager.Rooms = dbrooms
		}

		w.Header().Set("Content-Type", "application/json")

		roomsJson, err := json.Marshal(roomsManager)
		if err != nil {
			errRooms := mnc.Rooms{
				CurrentNumber: 0,
				MaxNumRooms:   0,
				MaxRoomSize:   0,
				Rooms:         nil,
			}
			errJson, _ := json.Marshal(errRooms)
			w.Write(errJson)
		} else {
			w.Write(roomsJson)
		}
	}
}
