package model

type Room struct {
	Name string `json:"name" binding:"required"`
}
