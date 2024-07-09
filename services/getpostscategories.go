package services

import (
	"database/sql"
	"log"

	"burakforum/database"
	"burakforum/models"
)

func GetPostCategories(CatID string) ([]models.Post, error) {
	var catModels []models.Category

	checkCatModels := "SELECT id, name, post_id FROM categories WHERE id = ?"
	rows, err := database.DB.Query(checkCatModels, CatID)
	if err != nil {
		log.Println("Veritabanında ID'ye ait kategori yok")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var catModel models.Category
		err := rows.Scan(&catModel.ID, &catModel.Name, &catModel.PostID)
		if err != nil {
			log.Println("Kategori bilgilerini okurken sorun oluştu")
			return nil, err
		}
		catModels = append(catModels, catModel)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	var posts []models.Post

	for _, catModel := range catModels {
		query := "SELECT id, title, user_id, content, created_at, likes FROM posts WHERE id = ?"
		var post models.Post
		err := database.DB.QueryRow(query, catModel.PostID).Scan(&post.ID, &post.Title, &post.UserID, &post.Content, &post.CreatedAt, &post.Likes)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Printf("Post ID %s bulunamadı\n", catModel.PostID)
				continue
			}
			log.Println("Post bilgilerini okurken sorun oluştu")
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
