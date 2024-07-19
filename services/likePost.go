package services

import (
	"database/sql"
	"errors"

	"burakforum/database"
	"burakforum/models"

	"github.com/google/uuid"
)

func AddLike(postID, userID string) error {
	likeID := uuid.New().String()

	insertQuery := "INSERT INTO likes_dislikes (id, post_id, user_id, is_like) VALUES (?, ?, ?, ?)"
	_, err := database.DB.Exec(insertQuery, likeID, postID, userID, true)
	if err != nil {
		return err
	}

	// Postun beğeni sayısını güncelle
	updatePostQuery := "UPDATE posts SET like_count = like_count + 1 WHERE id = ?"
	_, err = database.DB.Exec(updatePostQuery, postID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteLike(postID, userID string) error {
	var existLikeTable models.LikeDislike

	// Belirli bir postID ve userID'ye sahip beğeni kaydını kontrol et
	checkLikeQuery := "SELECT id, post_id, user_id, is_like FROM likes_dislikes WHERE post_id = ? AND user_id = ?"
	err := database.DB.QueryRow(checkLikeQuery, postID, userID).Scan(&existLikeTable.ID, &existLikeTable.PostID, &existLikeTable.UserID, &existLikeTable.IsLike)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("like does not exist")
		}
		return err
	}

	// Beğeni kaydını sil
	deleteQuery := "DELETE FROM likes_dislikes WHERE id = ?"
	_, err = database.DB.Exec(deleteQuery, existLikeTable.ID)
	if err != nil {
		return err
	}

	// Postun beğeni veya dislike sayısını güncelle
	var updatePostQuery string
	if existLikeTable.IsLike {
		updatePostQuery = "UPDATE posts SET like_count = like_count - 1 WHERE id = ?"
	} else {
		updatePostQuery = "UPDATE posts SET dislike_count = dislike_count - 1 WHERE id = ?"
	}
	_, err = database.DB.Exec(updatePostQuery, postID)
	if err != nil {
		return err
	}

	return nil
}

func LikePost(postID, userID string) error {
	var existliketable models.LikeDislike

	CheckLikeQuery := "SELECT id, post_id, user_id, is_like FROM likes_dislikes WHERE post_id = ? AND user_id = ?"
	err := database.DB.QueryRow(CheckLikeQuery, postID, userID).Scan(&existliketable.ID, &existliketable.PostID, &existliketable.UserID, &existliketable.IsLike)
	if err != nil && err == sql.ErrNoRows {
		AddLike(postID, userID)
	} else {
		DeleteLike(postID, userID)
	}
	return nil
}
