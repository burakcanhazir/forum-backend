package services

import (
	"database/sql"
	"fmt"
	"log"

	"burakforum/database"
	"burakforum/models"
)

func GetCategoriesPost(SelectPostLang string) ([]models.Post, error) {
	query := fmt.Sprintf("SELECT post_id FROM categories WHERE %s = '1'", SelectPostLang)
	log.Printf("Executing query: %s", query) // Debug log

	rows, err := database.DB.Query(query)
	if err != nil {
		log.Printf("Query execution failed: %v", err) // Debug log
		return nil, err
	}
	defer rows.Close()

	var postIDs []string

	for rows.Next() {
		var postID string
		if err := rows.Scan(&postID); err != nil {
			log.Printf("Row scan failed: %v", err) // Debug log
			return nil, err
		}
		postIDs = append(postIDs, postID)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Rows iteration failed: %v", err) // Debug log
		return nil, err
	}

	return GetPostCat(postIDs)
}

func GetPostCat(postIDs []string) ([]models.Post, error) {
	var posts []models.Post
	query := "SELECT id, user_id, title, content, created_at FROM posts WHERE id = ?"
	log.Printf("Fetching posts for IDs: %v", postIDs) // Debug log

	for _, postID := range postIDs {
		var post models.Post
		err := database.DB.QueryRow(query, postID).Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Printf("No post found for ID: %s", postID) // Debug log
				continue                                       // Eğer post bulunamazsa, loop devam etsin
			}
			log.Printf("QueryRow scan failed: %v", err) // Debug log
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
