package models

// Category struct defines the structure for category data
type Category struct {
	ID     string `json:"id"`     // Kategori ID'si
	Name   string `json:"name"`   // Kategori adı
	PostID string `json:"PostID"` // POST adı
}
