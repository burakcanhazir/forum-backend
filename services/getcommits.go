package services

import (
	"encoding/json"
	"log"

	"burakforum/database"
)

func GetCommits(PostID string) string {
	var CommitsPost []string

	// Veritabanından sorguyu çalıştır.
	Query := "SELECT content FROM commits WHERE post_id = ?"
	rows, err := database.DB.Query(Query, PostID)
	if err != nil {
		log.Println("Yorum bulma esnasında sorun:", err)
		return ""
	}
	defer rows.Close()

	// Satırları gezerek her bir yorumu `CommitsPost` slice'ine ekleyelim.
	for rows.Next() {
		var content string
		if err := rows.Scan(&content); err != nil {
			log.Println("Satır okunurken hata oluştu:", err)
			return ""
		}
		CommitsPost = append(CommitsPost, content)
	}

	// JSON verisini encode etme
	jsonData, err := json.Marshal(CommitsPost)
	if err != nil {
		log.Println("JSON dönüştürme esnasında hata oluştu:", err)
		return ""
	}

	return string(jsonData)
}
