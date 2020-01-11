package model

// Menu is representation of Restaurant Table
type Menu struct {
	ID    int    `json:"ID"`
	Name  string `json:"Nama"`
	Price int    `json:"Harga"`
}
