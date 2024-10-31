package services

import (
	"encoding/json"
	"log"

	"burakforum/database"
	"burakforum/models"
)

func GetCommits(PostID string) string {
	// Commit struct'larından oluşan bir dilim oluştur.
	var commits []models.Commit

	// Veritabanından ID ve content çekilecek.
	Query := "SELECT id, content FROM commits WHERE post_id = ?"
	rows, err := database.DB.Query(Query, PostID)
	if err != nil {
		log.Println("Yorum bulma esnasında sorun:", err)
		return ""
	}
	defer rows.Close()

	// Veritabanı satırlarını dolaş ve `commits` dilimine ekle.
	for rows.Next() {
		var commit models.Commit
		if err := rows.Scan(&commit.ID, &commit.Content); err != nil {
			log.Println("Satır okunurken hata oluştu:", err)
			return ""
		}
		commits = append(commits, commit)
	}

	// JSON verisine çevirme
	jsonData, err := json.Marshal(commits)
	if err != nil {
		log.Println("JSON dönüştürme esnasında hata oluştu:", err)
		return ""
	}
	return string(jsonData)
}
