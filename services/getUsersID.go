package services

import (
	"database/sql"
	"log"

	"burakforum/database"
	"burakforum/models"
)

func GetUserID(userID string) (*models.User, error) {
	var user models.User

	query := "SELECT id, name, email FROM users WHERE id = ?"
	err := database.DB.QueryRow(query, userID).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			// ID'ye sahip kullanıcı bulunamadı
			return nil, nil
		}
		log.Printf("Error querying user by ID: %v", err)
		return nil, err
	}

	return &user, nil
}
