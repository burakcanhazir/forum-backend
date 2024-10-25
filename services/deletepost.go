package services

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"burakforum/database"
	"burakforum/models"
)

var ErrNoPostFound = errors.New("no post found with the given ID")

func DeletePost(postID string) error {
	fmt.Println(postID)
	check := "SELECT id, title, content, user_id, created_at FROM posts WHERE id = ?"
	var exp models.Post

	err := database.DB.QueryRow(check, postID).Scan(&exp.ID, &exp.Title, &exp.Content, &exp.UserID, &exp.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("post yok: %v", err)
			return ErrNoPostFound
		}
		log.Printf("post sorgusunda hata oluştu: %v", err)
		return err
	}

	query2 := "DELETE FROM posts WHERE id = ?"
	_, err = database.DB.Exec(query2, exp.ID)
	if err != nil {
		log.Printf("post silinirken hata oluştu: %v", err)
		return err
	}
	return nil
}
