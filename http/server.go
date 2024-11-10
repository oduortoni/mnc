package http

import (
	"net/http"

	"mnc/mnc"
)

func Server(address string, roomsManager *mnc.Rooms) {
	http.HandleFunc("/", Index)
	http.ListenAndServe(address, nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	// protocol
}
