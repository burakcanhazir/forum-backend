package services

import (
	"log"

	"burakforum/database"
	"burakforum/models"

	"github.com/google/uuid"
)

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
