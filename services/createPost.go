package services

import (
	"burakforum/database"
	"burakforum/models"
)

func CreatePost(post *models.Post) error {
	query := "INSERT INTO posts (id, title, user_id, content, created_at) VALUES (?, ?, ?, ?, ?)"
	_, err := database.DB.Exec(query, post.ID, post.Title, post.UserID, post.Content, post.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
