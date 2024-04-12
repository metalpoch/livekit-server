package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/metalpoch/livekit-server/pkg/streaming"
)

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "name is required")
		return
	}

	room, err := streaming.CreateRoom(name)

	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&room)
}
