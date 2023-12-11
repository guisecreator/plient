package models

import "time"

type BoardCounts struct {
	Pins          int32 `json:"pins"`
	Collaborators int32 `json:"collaborators"`
	Followers     int32 `json:"followers"`
}

type Boards struct {
	Items    []Board `json:"items"`
	Bookmark string  `json:"bookmark"`
}

type Board struct {
	Id          string       `json:"id"`
	Name        string       `json:"name"`
	Url         string       `json:"url"`
	Description string       `json:"description"`
	Creator     Creator      `json:"creator"`
	CreatedAt   time.Time `json:"created_at"`
	Counts      BoardCounts  `json:"counts"`
	Image       Images       `json:"image"`
	Privacy     string       `json:"privacy"`
}