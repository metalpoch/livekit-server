package routes

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/metalpoch/livekit-server/handler"
)

func GenerateTokenToJoin(w http.ResponseWriter, r *http.Request) {
	var guess_name string = uuid.New().String()

	room := r.URL.Query().Get("room")
	if room == "" {
		fmt.Fprint(w, "room param is required")
		return
	}

	token, err := handler.Token(room, guess_name)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	fmt.Fprint(w, token)

}
