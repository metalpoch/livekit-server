package routes

import (
	"encoding/json"
	"net/http"

	"github.com/metalpoch/livekit-server/handler"
	"github.com/metalpoch/livekit-server/model"
)

func Room(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		handler.GetRooms(w, r)
	case http.MethodPost:
		handler.CreateRoom(w, r)
	case http.MethodDelete:
		handler.DeleteRoom(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(&model.Exception{
			Error: "method not allowed",
		})
	}
}
