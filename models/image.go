package models

type Image struct {
	Url string `json:"url"`
}

type Media struct {
	Images Images `json:"images"`
}

type Images struct {
	Res150x150 Image `json:"150x150"`
	Res400x300 Image `json:"400x300"`
	Res600x    Image `json:"600x"`
	Res1200x   Image `json:"1200x"`
}