package services

import (
	"database/sql"
	"errors"
	"log"

	"burakforum/database"
	"burakforum/models"
)

func CreateUser(user *models.User) error {
	// Kullanıcının adı zaten var mı kontrol et
	var existingUser models.User

	checkQuery := "SELECT id, name, email FROM users WHERE name = ?"
	err := database.DB.QueryRow(checkQuery, user.Name).Scan(&existingUser.ID, &existingUser.Name, &existingUser.Email)

	if err != nil && err != sql.ErrNoRows {
		log.Printf("Error checking for existing user: %v", err)
		return err
	}
	if existingUser.Name != "" {
		return errors.New("user with the same name already exists")
	}

	// Yeni kullanıcıyı ekle
	insertQuery := "INSERT INTO users (id, name, email) VALUES (?, ?, ?)"
	_, err = database.DB.Exec(insertQuery, user.ID, user.Name, user.Email)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return err
	}
	return nil
}
