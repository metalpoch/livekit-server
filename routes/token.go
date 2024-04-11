package routes

import (
	"encoding/json"
	"net/http"

	"github.com/metalpoch/livekit-server/handler"
	"github.com/metalpoch/livekit-server/model"
)

func GetToken(w http.ResponseWriter, r *http.Request) {
	var room string
	var identity string

	queryParams := r.URL.Query()
	w.Header().Set("Content-Type", "application/json")

	if room = queryParams.Get("room"); room == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&model.Exception{
			Error: "room param is required",
		})
		return
	}

	if identity = queryParams.Get("identity"); identity == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&model.Exception{
			Error: "identity param is required",
		})
		return
	}
	token, err := handler.GetJoinToken(room, identity)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(&model.Token{
		Token:    token,
		Room:     room,
		Identity: identity,
	})
}
