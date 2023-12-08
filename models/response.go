package models

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Type    string      `json:"type"`
	Page    Page        `json:"page"`
}

type Page struct {
	Cursor string `json:"cursor"`
	Next   string `json:"next"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
