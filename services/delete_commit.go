package services

import (
	"database/sql"
	"log"

	"forumbackend/database"
)

func DeleteCommit(userID, postID, commitID string) error {
	var fieldedUser string
	var fieldedPost string

	Query := "SELECT user_id, post_id FROM commits WHERE id = ?"
	err := database.DB.QueryRow(Query, commitID).Scan(&fieldedUser, &fieldedPost)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("COMMİT YOK")
			return err
		}
		log.Printf("Commit aranırken hata oluştu: %v", err)
		return err
	}

	var ispostU string

	Query2 := "SELECT user_id FROM posts WHERE id = ?"
	err = database.DB.QueryRow(Query2, fieldedPost).Scan(&ispostU)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Postun sahibi bulunamadı")
			return err
		}
		log.Printf("Postun sahibi aranırken hata oluştu: %v", err)
		return err
	}
	if userID == ispostU || userID == fieldedUser {
		Query3 := "DELETE FROM commits WHERE id = ?"
		_, err = database.DB.Exec(Query3, commitID)
		if err != nil {
			log.Printf("Commit silinirken hata oluştu: %v", err)
			return err
		}
		log.Println("Commit başarıyla silindi")
	} else {
		log.Printf("Postun sahibi aranırken hata oluştu: %v", err)
		return err
	}
	return nil
}
