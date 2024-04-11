package handler

import (
	"encoding/json"
	"net/http"

	"github.com/metalpoch/livekit-server/model"
	"github.com/metalpoch/livekit-server/pkg/streaming"
)

func GetRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := streaming.ListRooms()

	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(&model.Exception{
			Error: err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(&rooms)
}
