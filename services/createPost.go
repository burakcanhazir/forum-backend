package services

import (
	"encoding/json"
	"fmt"
	"strings"

	"burakforum/database"
	"burakforum/models"
	"burakforum/utils"

	"github.com/google/uuid"
)

// generateUUID creates a new UUID string
func generateUUID() string {
	return uuid.New().String()
}

func CreatePost(post *models.Post) error {
	// Convert the category slice to JSON
	categoryJSON, err := json.Marshal(post.Category)
	if err != nil {
		return err
	}

	query := "INSERT INTO posts (id, title, user_id, content, category, created_at) VALUES (?, ?, ?, ?, ?, ?)"
	_, err = database.DB.Exec(query, post.ID, post.Title, post.UserID, post.Content, string(categoryJSON), post.CreatedAt)
	if err != nil {
		return err
	}
	// Initialize all category columns to 0
	categoryValues := make(map[string]int)
	for key := range utils.CategoryColumns {
		categoryValues[key] = 0
	}

	// Set the corresponding category columns to 1
	for _, category := range post.Category {
		columnName, exists := utils.CategoryColumns[category]
		if exists {
			categoryValues[columnName] = 1
		}
	}

	// Create a slice to hold the column names and values
	var columns []string
	var values []interface{}
	columns = append(columns, "id", "name", "post_id")
	values = append(values, generateUUID(), "post_category", post.ID)
	for key, value := range categoryValues {
		columns = append(columns, key)
		values = append(values, value)
	}

	// Construct the SQL query
	insertQuery := fmt.Sprintf("INSERT INTO categories (%s) VALUES (%s)",
		strings.Join(columns, ", "),
		strings.TrimRight(strings.Repeat("?, ", len(values)), ", "))

	_, err = database.DB.Exec(insertQuery, values...)
	if err != nil {
		return err
	}

	return nil
}
