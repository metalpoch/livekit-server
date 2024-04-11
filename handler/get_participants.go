package handler

import (
	"encoding/json"
	"net/http"

	"github.com/metalpoch/livekit-server/model"
	"github.com/metalpoch/livekit-server/pkg/streaming"
)

func GetParticipants(w http.ResponseWriter, r *http.Request) {
	var room string
	queryParams := r.URL.Query()

	if room = queryParams.Get("room"); room == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&model.Exception{
			Error: "room param is required",
		})
		return
	}

	participants, err := streaming.GetParticipants(room)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(&model.Exception{
			Error: err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(&participants)
}
