package routes

import (
	"encoding/json"
	"net/http"

	"github.com/metalpoch/livekit-server/handler"
	"github.com/metalpoch/livekit-server/model"
)

func Participant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		handler.GetParticipants(w, r)
	case http.MethodPut:
		handler.MuteOrUnmuteParticipantsTracks(w, r)
	case http.MethodDelete:
		handler.RemoveParticipant(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(&model.Exception{
			Error: "method not allowed",
		})
	}
}
