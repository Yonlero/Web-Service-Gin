package entities

import "github.com/google/uuid"

// -----------------------------Structs-------------------------
type Album struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Artist string    `json:"artist"`
	Price  float64   `json:"price"`
}
