package services

import (
	"burakforum/database"
	"burakforum/models"
	"errors"
	"log"
)

func FieldPosts(userID string) ([]string, error) {
	var postIDs []string

	query := "SELECT post_id FROM likes_dislikes WHERE user_id = ? AND is_like = 1"
	rows, err := database.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var postID string
		if err := rows.Scan(&postID); err != nil {
			return nil, err
		}
		postIDs = append(postIDs, postID)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return postIDs, nil
}

func UsersLikesPost(userID string) ([]models.Post, error) {
	var fieldedposts []models.Post

	postslices, err := FieldPosts(userID)
	if err != nil {
		log.Printf("BEĞENİLEN POSTLAR ÇEKİLİRKEN HATA OLUŞTU: %v", err)
		return nil, err
	}

	if len(postslices) == 0 {
		log.Printf("BEĞENİLEN POST YOK")
		return nil, errors.New("no liked posts found")
	}

	for _, r := range postslices {
		var fieldpost models.Post
		checkposts := "SELECT id, title, user_id, content, created_at FROM posts WHERE id = ?"
		err := database.DB.QueryRow(checkposts, r).Scan(&fieldpost.ID, &fieldpost.Title, &fieldpost.UserID, &fieldpost.Content, &fieldpost.CreatedAt)
		if err != nil {
			log.Printf("databaseden post çekilirken hata oluştu: %v", err)
			continue
		}
		fieldedposts = append(fieldedposts, fieldpost)
	}

	return fieldedposts, nil
}
