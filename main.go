package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/metalpoch/livekit-server/routes"
)

func checkEnvVar() {
	variables := [3]string{"LIVEKIT_HOST_URL", "LIVEKIT_API_KEY", "LIVEKIT_API_SECRET"}
	length := len(variables)
	for i := 0; i < length; i++ {
		if variables[i] == "" {
			fmt.Println(variables[i], "variable was requried")
			os.Exit(0)
		}
	}
}

func main() {
	checkEnvVar()
	mux := http.NewServeMux()
	mux.HandleFunc("/api/token", routes.GetToken)
	mux.HandleFunc("/api/room", routes.Room)
	mux.HandleFunc("/api/participants", routes.Participant)

	log.Fatal(http.ListenAndServe(":8080", mux))
}