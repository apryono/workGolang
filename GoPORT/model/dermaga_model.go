package model

// Dermaga struct
type Dermaga struct {
	ID       int    `json:"Dock_ID,omitempty"`
	Kode     string `json:"Kode_ID"`
	Status   string `json:"Status"`
	IsDelete int    `json:"Is_Delete,omitempty"`
}
