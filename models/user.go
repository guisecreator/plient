package models

import "time"

type User struct {
	Id          string       `json:"id"`
	Username    string       `json:"username"`
	FirstName   string       `json:"first_name"`
	LastName    string       `json:"last_name"`
	Bio         string       `json:"bio"`
	AccountType string       `json:"account_type"`
	Url         string       `json:"account_type"`
	CreatedAt   time.Time    `json:"created_at"`
	Counts      UserCounts   `json:"counts"`
	Image       Images       `json:"image"`
}

type Creator struct {
	Url       string `json:"url"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Id        string `json:"id"`
}

type UserCounts struct {
	Pins 		int `json:"pins"`
	Boards 		int `json:"boards"`
	Followers 	int `json:"followers"`
	Following 	int `json:"following"`
	Likes 		int `json:"likes"`
}