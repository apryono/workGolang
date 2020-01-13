package model

// Kapal struct is data type for all table
type Kapal struct {
	ID       int    `json:"ID_Kapal,omitempty"`
	Kode     string `json:"Kode_ID"`
	Muatan   int    `json:"Muatan"`
	Status   string `json:"Status"`
	IsDelete int    `json:"Is_Delete,omitempty"`
}
