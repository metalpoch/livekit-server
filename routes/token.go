package routes

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/metalpoch/livekit-server/handler"
)

func GenerateToken(w http.ResponseWriter, r *http.Request) {
	var room string = uuid.New().String()

	username := r.URL.Query().Get("username")
	if username == "" {
		fmt.Fprint(w, "username param is required")
		return
	}

	token, err := handler.Token(room, username)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	fmt.Fprint(w, token)

}
