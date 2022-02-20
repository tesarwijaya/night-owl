package model

type PlayerModel struct {
	Name   string `json:"name,omitempty"`
	TeamID int64  `json:"teamId,omitempty"`
}
