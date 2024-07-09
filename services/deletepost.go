package services

import (
	"errors"

	"burakforum/database"
)

func DeletePost(postID string) error {
	check := "DELETE FROM posts WHERE id = ?"

	stmt, err := database.DB.Prepare(check)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(postID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no post found with the given ID")
	}

	return nil
}
