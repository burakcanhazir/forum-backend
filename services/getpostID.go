package services

import (
	"database/sql"
	"log"

	"burakforum/database"
	"burakforum/models"
)

func GetpostID(postID string) (*models.Post, error) {
	var fieldpost models.Post
	checkpost := "SELECT * FROM posts WHERE id = ?"

	err := database.DB.QueryRow(checkpost, postID).Scan(&fieldpost.ID, &fieldpost.Title, &fieldpost.UserID, &fieldpost.Content, &fieldpost.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("postID bulunamadı: ", postID)
			return nil, nil
		}
		log.Println("Query error:", err)
		return nil, err
	}

	return &fieldpost, nil
}
