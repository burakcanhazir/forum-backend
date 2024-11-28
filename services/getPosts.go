package services

import (
	"forumbackend/database"
	"forumbackend/models"

	_ "github.com/mattn/go-sqlite3"
)

// GetPosts fonksiyonu, veri tabanından postları döndürür
func GetPosts() ([]models.Post, error) {
	rows, err := database.DB.Query("SELECT id, user_id, created_at, title, content FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.UserID, &post.CreatedAt, &post.Title, &post.Content); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
