package rss

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) Store {
	return Store{
		db: db,
	}
}

func (s Store) CreateFeed(ctx context.Context, feed Feed) (Feed, error) {
	result := s.db.Create(&feed)
	if result.Error != nil {
		return Feed{}, result.Error
	}
	return feed, nil
}

func (s Store) QueryByID(ctx context.Context, id int) (Feed, error) {
	var feed Feed
	res := s.db.Where(&Feed{
		Model:       gorm.Model{ID: uint(id)},
		Link:        "",
		Title:       "",
		Description: sql.NullString{},
	}).First(&feed)
	if res.Error != nil {
		return Feed{}, res.Error
	}
	return feed, nil
}

func (s Store) GetFeedsWithLastPublishedTime() []Feed {
	var feeds []Feed
	query := s.db.Model(&Feed{}).Select("feeds.*, max(posts.published_at) as last_published_at").
		Joins("left join posts on posts.feed_id = feeds.id").Group("feeds.id")
	query.Find(&feeds)
	return feeds
}

func (s Store) CreatePost(ctx context.Context, p Post) (Post, error) {
	result := s.db.Create(&p)
	if result.Error != nil {
		return Post{}, result.Error
	}
	return p, nil
}
