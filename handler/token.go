package handler

import (
	"time"

	"github.com/livekit/protocol/auth"
)

func Token(room, identity string) (string, error) {
	apiKey := "devkey"
	apiSecret := "secret"

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
