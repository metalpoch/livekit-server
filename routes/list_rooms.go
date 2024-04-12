package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/metalpoch/livekit-server/pkg/streaming"
)

func ListAvailableRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := streaming.ListRooms()

	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprint(w, rooms)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&rooms)

}
