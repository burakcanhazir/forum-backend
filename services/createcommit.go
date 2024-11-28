package services

import (
	"database/sql"
	"log"

	"forumbackend/database"
	"forumbackend/models"

	"github.com/google/uuid"
)

func TruePostID(postID string) error {
	var istrue models.Post

	Query := "SELECT id, title, user_id, content FROM posts WHERE id = ?"
	err := database.DB.QueryRow(Query, postID).Scan(&istrue.ID, &istrue.Title, &istrue.UserID, &istrue.Content)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Post bulunamadı")
			return err
		}
		log.Printf("Post aranırken hata oluştu: %v", err)
		return err
	}

	// Post bulundu, işleme devam edilebilir
	log.Printf("Post bulundu: %+v", istrue)
	return nil
}

func CreateCommit(Commits *models.Commit) error {
	Commits.ID = uuid.New().String()
	Query := "INSERT INTO commits (id , user_id, post_id, content) VALUES (?,?,?,?)"
	_, err := database.DB.Exec(Query, Commits.ID, Commits.UserID, Commits.PostID, Commits.Content)
	if err != nil {
		log.Printf("dataya atılırken sorun oluştu")
		return err
	}
	return nil
}
