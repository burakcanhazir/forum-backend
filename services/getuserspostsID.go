package services

import (
	"log"
	"strings"

	"forumbackend/database"
	"forumbackend/models"
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
		var category string // Category'yi string olarak alacağız
		err := row.Scan(&post.ID, &post.Title, &post.UserID, &post.Content, &category, &post.CreatedAt, &post.Likes)
		if err != nil {
			log.Fatal(err)
		}

		// Category'yi []string'e dönüştürüyoruz
		post.Category = strings.Split(category, ",") // Kategoriler virgülle ayrılmışsa
		posts = append(posts, post)
	}
	return posts, nil
}
