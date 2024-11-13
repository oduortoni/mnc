package http

import (
	"net/http"

	"mnc/mnc"
)

func Server(address string, roomsManager *mnc.Rooms) {
	http.HandleFunc("/static/", Static)
	http.HandleFunc("/explore", ExploreAll(roomsManager))
	http.HandleFunc("/explore/", ExploreOne(roomsManager))
	http.HandleFunc("/createroom", CreateRoom(roomsManager))
	http.HandleFunc("/messages/save", SaveMessage(roomsManager))
	http.HandleFunc("/error", Error)
	http.HandleFunc("/", Index)
	http.ListenAndServe(address, nil)
}
