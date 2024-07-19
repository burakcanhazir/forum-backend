package services

import (
	"burakforum/database"
	"burakforum/models"
)

func CreatePost(post *models.Post) error {
	query := "INSERT INTO posts (title, user_id, content, created_at) VALUES (?, ?, ?, ?)"
	result, err := database.DB.Exec(query, post.Title, post.UserID, post.Content, post.CreatedAt)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	post.ID = int(id)
	return nil
}
