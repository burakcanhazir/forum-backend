package services

import (
	"database/sql"
	"encoding/json"
	"log"

	"forumbackend/database"
	"forumbackend/models"
)

func GetpostID(postID string) (*models.Post, error) {
	log.Printf("burası services back")
	var fieldpost models.Post
	var categoryJSON string

	checkpost := "SELECT * FROM posts WHERE id = ?"

	err := database.DB.QueryRow(checkpost, postID).Scan(
		&fieldpost.ID,
		&fieldpost.Title,
		&fieldpost.UserID,
		&fieldpost.Content,
		&fieldpost.CreatedAt, // created_at TEXT olarak saklandığı için string
		&categoryJSON,        // category JSON formatında TEXT olarak saklanıyor
		&fieldpost.Likes,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("postID bulunamadı: ", postID)
			return nil, nil
		}
		log.Println("Query error:", err)
		return nil, err
	}

	// JSON olarak saklanan category verisini []string'e dönüştür
	if err := json.Unmarshal([]byte(categoryJSON), &fieldpost.Category); err != nil {
		log.Println("Failed to parse category JSON:", err)
		return nil, err
	}

	return &fieldpost, nil
}
