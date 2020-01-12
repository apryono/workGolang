package model

// ResponseWrapper represent standard response message
type ResponseWrapper struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Success represent standard response message
type Success struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Error struct
type Error struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
