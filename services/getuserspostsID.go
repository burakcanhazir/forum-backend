package services

import (
	"log"

	"burakforum/database"
	"burakforum/models"
)

func GetUsersPostsID(userID string) ([]models.Post, error) {
	var posts []models.Post
	var post models.Post

	checkuserspost := "SELECT * FROM posts WHERE user_id = ?"
	row, err := database.DB.Query(checkuserspost, userID)
	if err != nil {
		log.Fatal(err)
	}

	for row.Next() {
		err := row.Scan(&post.ID, &post.Title, &post.UserID, &post.Content, &post.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, post)
	}
	return posts, nil
}
