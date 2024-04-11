package handler

import (
	"encoding/json"
	"net/http"

	"github.com/metalpoch/livekit-server/model"
	"github.com/metalpoch/livekit-server/pkg/streaming"
)

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	var body model.Room
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&model.Exception{
			Error: err.Error(),
		})
		return
	}

	if body.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&model.Exception{
			Error: "name is required",
		})
		return
	}

	room, err := streaming.CreateRoom(body.Name)

	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(&model.Exception{
			Error: err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(&room)
}
