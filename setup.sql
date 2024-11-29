package database

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func createTables() {
	blackListToken := `
	CREATE TABLE IF NOT EXISTS blacklist (
		token TEXT PRIMARY KEY,
        tokenstring TEXT
	);`

	userTable := `
    CREATE TABLE IF NOT EXISTS users (
        id TEXT PRIMARY KEY,
        name TEXT,
        email TEXT,
        password TEXT
    );`

	postTable := `
    CREATE TABLE IF NOT EXISTS posts (
        id TEXT PRIMARY KEY,
        title TEXT,
        user_id TEXT,
        content TEXT,
        created_at TEXT,
        category TEXT,
        like_count INTEGER DEFAULT 0,
        FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
    );`

	likeDislikeTable := `
    CREATE TABLE IF NOT EXISTS likes_dislikes (
        id TEXT PRIMARY KEY,
        post_id TEXT,
        user_id TEXT,
        is_like BOOLEAN,
        FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE,
        FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
    );`

	createCommit := `
    CREATE TABLE IF NOT EXISTS commits (
        id TEXT PRIMARY KEY,
        user_id TEXT,
        post_id TEXT,
        content TEXT,
        FOREIGN KEY(user_id) REFERENCES users(id),
        FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE
    );`

	categoryTable := `
    CREATE TABLE IF NOT EXISTS categories (
        id TEXT PRIMARY KEY,
        name TEXT,
        post_id TEXT,
        go TEXT,
        php TEXT,
        python TEXT,
        c TEXT,
        csharp TEXT,
        cplus TEXT,
        rust TEXT,
        java TEXT,
        javascript TEXT,
        html TEXT,
        css TEXT,
        react TEXT,
        flutter TEXT,
        assembly TEXT,
        perl TEXT,
        swift TEXT,
        other TEXT,
        FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE
    );`

	_, err := DB.Exec(userTable)
	if err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}
	_, err = DB.Exec(blackListToken)
	if err != nil {
		log.Fatal(err)
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
