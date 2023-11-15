package database

import (
	"log"
)

type Post struct {
	Id        int
	Content   string
	CreatedAt string
}

func GetPosts() []Post {
	posts := []Post{}
	rows, _ := DB.Query("SELECT content, created_at FROM post")

	// Placeholder variables for query result
	var content string
	var created_at string
	for rows.Next() {
		err := rows.Scan(&content, &created_at)
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, Post{Content: content, CreatedAt: created_at})
	}
	defer rows.Close()
	return posts
}
