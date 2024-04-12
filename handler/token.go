package handler

import (
	"os"
	"time"

	"github.com/livekit/protocol/auth"
)

func Token(room, identity string) (string, error) {
	apiKey := os.Getenv("LIVEKIT_API_KEY")
	apiSecret := os.Getenv("LIVEKIT_API_SECRET")

	at := auth.NewAccessToken(apiKey, apiSecret)
	grant := &auth.VideoGrant{
		RoomJoin: true,
		Room:     room,
	}
	at.AddGrant(grant).
		SetName(identity).
		SetIdentity(identity).
		SetValidFor(48 * time.Hour)

	return at.ToJWT()
}
