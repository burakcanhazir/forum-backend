package models

import (
	"database/sql"
)

// Post struct defines the structure for post data
type Post struct {
	ID        string        `json:"id"` // Post ID'si
	Title     string        `json:"title"`
	UserID    string        `json:"user_id"` // Postu oluşturan kullanıcının ID'si
	Content   string        `json:"content"` // Post içeriği
	Category  []string      `json:"category"`
	CreatedAt string        `json:"created_at"` // Postun oluşturulma tarihi
	Likes     sql.NullInt64 `json:"likes"`      // Nullable integer type
}
