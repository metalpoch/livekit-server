package handler

import (
	"encoding/json"
	"net/http"

	"github.com/metalpoch/livekit-server/model"
	"github.com/metalpoch/livekit-server/pkg/streaming"
)

func RemoveParticipant(w http.ResponseWriter, r *http.Request) {
	var body model.Participant
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(&model.Exception{
			Error: err.Error(),
		})
		return
	}

	err = streaming.RemoveParticipant(body.Room, body.Identity)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(&model.Exception{
			Error: err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(true)
}
