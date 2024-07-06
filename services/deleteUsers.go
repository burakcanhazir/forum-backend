package services

import (
	"database/sql"
	"log"

	"burakforum/database"
	"burakforum/models"
)

func DeleteUsers(user *models.User) error {
	var deleteUs models.User

	checkQuery := "SELECT id, name, email FROM users WHERE name = ?"
	err := database.DB.QueryRow(checkQuery, user.Name).Scan(&deleteUs.ID, &deleteUs.Name, &deleteUs.Email)
	if err == sql.ErrNoRows {
		log.Printf("kullanıcı yok: %v", err)
		return err
	} else if err != nil {
		log.Printf("Veritabanı hatası: %v", err)
		return err
	}

	deleteuserID := "DELETE FROM users WHERE name = ?"
	_, err = database.DB.Exec(deleteuserID, user.Name)
	if err != nil {
		log.Printf("Error NOT DELETE USER: %v", err)
		return err
	}
	return nil
}
