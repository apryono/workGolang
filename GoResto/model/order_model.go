package model

// Order struct
type Order struct {
	MejaID  int           `json:"MejaID"`
	Pesanan []DetailOrder `json:"Pesanan"`
}

// DetailOrder struct
type DetailOrder struct {
	MenuID int    `json:"MenuID"`
	Qty    int    `json:"Qty"`
	Notes  string `json:"Notes"`
}
