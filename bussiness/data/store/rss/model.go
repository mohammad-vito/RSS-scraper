package rss

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Feed struct {
	gorm.Model
	Link            string         `gorm:"link" json:"link"`
	Title           string         `json:"title"`
	Description     sql.NullString `gorm:"type:text" json:"description"`
	Posts           []Post
	LastPublishedAt sql.NullTime `gorm:"<-:false"`
}

type Post struct {
	gorm.Model
	Link        string         `gorm:"link" json:"link"`
	Title       string         `json:"title"`
	Author      string         `json:"author"`
	Content     sql.NullString `gorm:"type:text" json:"content"`
	FeedID      uint           `json:"feed_id"`
	PublishedAt sql.NullTime   `json:"published_at"`
}

func NewPost(link, title, author, content string, fID uint, publishedAt time.Time) Post {
	//p := Post{}
	return Post{
		Link:   link,
		Title:  title,
		Author: author,
		FeedID: fID,
		Content: sql.NullString{
			String: content,
			Valid:  true,
		},
		PublishedAt: sql.NullTime{
			Time:  publishedAt,
			Valid: true,
		},
	}
}
