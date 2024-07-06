package database

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func createTables() {
	userTable := `
    CREATE TABLE IF NOT EXISTS users (
        id TEXT PRIMARY KEY,
        name TEXT,
        email TEXT
    );`

	postTable := `
    CREATE TABLE IF NOT EXISTS posts (
        id TEXT PRIMARY KEY,
        title TEXT,
        user_id TEXT,
        content TEXT,
        created_at TEXT,
        FOREIGN KEY(user_id) REFERENCES users(id)
    );`

	likeDislikeTable := `
    CREATE TABLE IF NOT EXISTS likes_dislikes (
        id TEXT PRIMARY KEY,
        post_id TEXT,
        user_id TEXT,
        is_like BOOLEAN,
        FOREIGN KEY(post_id) REFERENCES posts(id),
        FOREIGN KEY(user_id) REFERENCES users(id)
    );`

	createCommit := `
    CREATE TABLE IF NOT EXISTS commits (
        id TEXT PRIMARY KEY,
        user_id TEXT,
        post_id TEXT,
        content TEXT,
        created_at TEXT,
        FOREIGN KEY(user_id) REFERENCES users(id),
        FOREIGN KEY(post_id) REFERENCES posts(id)
    );`

	categoryTable := `
    CREATE TABLE IF NOT EXISTS categories (
        id TEXT PRIMARY KEY,
        name TEXT
    );`

	_, err := DB.Exec(userTable)
	if err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}

	_, err = DB.Exec(postTable)
	if err != nil {
		log.Fatalf("Failed to create posts table: %v", err)
	}

	_, err = DB.Exec(likeDislikeTable)
	if err != nil {
		log.Fatalf("Failed to create likes_dislikes table: %v", err)
	}
	_, err = DB.Exec(createCommit)
	if err != nil {
		log.Fatalf("failed to create commits: %v", err)
	}

	_, err = DB.Exec(categoryTable)
	if err != nil {
		log.Fatalf("Failed to create categories table: %v", err)
	}
}
