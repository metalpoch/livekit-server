package model

type Participant struct {
	Room     string `json:"room" binding:"required"`
	Identity string `json:"identity" binding:"required"`
}

type MuteOrUnmuteParticipant struct {
	Room     string `json:"room" binding:"required"`
	Identity string `json:"identity" binding:"required"`
	TrackSid string `json:"track_sid" binding:"required"`
	Muted    bool   `json:"muted" binding:"required"`
}
