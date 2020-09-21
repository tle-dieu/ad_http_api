package model

type Ad struct {
	Brand   string   `json:"brand"`
	Model   string   `json:"model"`
	Price   int      `json:"price"`
	Options *Options `json:"options"`
}

type Options struct {
	Bluetooth bool `json:"bluetooth"`
	Gps       bool `json:"gps"`
}
