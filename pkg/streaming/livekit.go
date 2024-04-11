package streaming

import (
	"context"
	"os"

	"github.com/livekit/protocol/livekit"
	lksdk "github.com/livekit/server-sdk-go/v2"
)

func getRoomClient() *lksdk.RoomServiceClient {
	return lksdk.NewRoomServiceClient(
		os.Getenv("LIVEKIT_HOST_URL"),
		os.Getenv("LIVEKIT_API_KEY"),
		os.Getenv("LIVEKIT_API_SECRET"),
	)
}

func CreateRoom(roomName string) (*livekit.Room, error) {
	room, err := getRoomClient().CreateRoom(
		context.Background(), &livekit.CreateRoomRequest{
			Name: roomName,
		})
	return room, err
}

func ListRooms() (*livekit.ListRoomsResponse, error) {
	res, err := getRoomClient().ListRooms(context.Background(), &livekit.ListRoomsRequest{})
	return res, err
}

func DeleteRoom(roomName string) error {
	_, err := getRoomClient().DeleteRoom(context.Background(), &livekit.DeleteRoomRequest{
		Room: roomName,
	})
	return err
}

func GetParticipants(roomName string) (*livekit.ListParticipantsResponse, error) {
	res, err := getRoomClient().ListParticipants(context.Background(), &livekit.ListParticipantsRequest{
		Room: roomName,
	})
	return res, err
}

func RemoveParticipant(roomName, identity string) error {
	_, err := getRoomClient().RemoveParticipant(context.Background(), &livekit.RoomParticipantIdentity{
		Room:     roomName,
		Identity: identity,
	})
	return err
}

func MuteOrUnmuteParticipantsTracks(roomName, identity, track_sid string, muted bool) error {
	_, err := getRoomClient().MutePublishedTrack(context.Background(), &livekit.MuteRoomTrackRequest{
		Room:     roomName,
		Identity: identity,
		TrackSid: track_sid,
		Muted:    muted,
	})
	return err
}
