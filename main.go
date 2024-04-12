package main

import (
	"log"
	"net/http"

	"github.com/metalpoch/livekit-server/routes"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/generate-token", routes.GenerateToken)
	mux.HandleFunc("/generate-token-to-join", routes.GenerateTokenToJoin)
	mux.HandleFunc("/get-available-rooms", routes.ListAvailableRooms)
	mux.HandleFunc("/create-room", routes.CreateRoom)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
