package models

type PinData struct {
	Id            string `json:"id"`
	Title         string `json:"title"`
	DominantColor string `json:"dominant_color"`
	Media         Media  `json:"media"`
	BoardId       string `json:"board_id"`
	AltText       string `json:"alt_text"`
}

type PinsData struct {
	Items []PinData `json:"pins"`
	Image string    `json:"image"`
}