package http

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"

	"mnc/mnc"
)

func Server(address string, roomsManager *mnc.Rooms) {
	http.HandleFunc("/static/", Static)
	http.HandleFunc("/explore", Explore)
	http.HandleFunc("/error", Error)
	http.HandleFunc("/", Index)
	http.ListenAndServe(address, nil)
}

func Explore(w http.ResponseWriter, r *http.Request) {
	rooms := mnc.NewRooms(4, 5)
	rooms.CreateRoom("bounty", 4)
	rooms.CreateRoom("chocolate", 8)
	rooms.CreateRoom("neon", 9)

	w.Header().Set("Content-Type", "application/json")

	roomsJson, err := json.Marshal(rooms)
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

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/error", http.StatusFound)
		return
	}
	t := template.Must(template.ParseFiles("templates/index.html"))
	t.Execute(w, nil)
}

func Static(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	file := "." + url

	stat, err := os.Stat(file)
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusFound)
		return
	}

	if stat.IsDir() {
		http.Redirect(w, r, "/error", http.StatusFound)
		return
	}

	http.ServeFile(w, r, file)
}

func Error(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/errors.html"))
	t.Execute(w, nil)
}
