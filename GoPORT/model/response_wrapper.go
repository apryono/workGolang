package model

// Success represent standard response message
type Success struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Error struct
type Error struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
