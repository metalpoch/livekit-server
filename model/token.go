package model

type Token struct {
	Token    string `json:"token"`
	Room     string `json:"room" binding:"required"`
	Identity string `json:"identity" binding:"required"`
}
