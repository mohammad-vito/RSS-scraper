package rss

import (
	"RssReader/bussiness/data/store/rss"
	"context"
	"fmt"
	"gorm.io/gorm"
)

type Core struct {
	rss rss.Store
}

func NewCore(db *gorm.DB) Core {
	return Core{rss: rss.NewStore(db)}
}

func (c Core) QueryByID(ctx context.Context, userID int) (rss.Feed, error) {
	feed, err := c.rss.QueryByID(ctx, userID)
	if err != nil {
		return rss.Feed{}, fmt.Errorf("query: %w", err)
	}
	return feed, nil
}

func (c Core) CreateFeed(ctx context.Context, feed rss.Feed) (rss.Feed, error) {
	feed, err := c.rss.CreateFeed(ctx, feed)
	if err != nil {
		return rss.Feed{}, fmt.Errorf("create: %w", err)
	}
	return feed, err
}

func (c Core) CreatePost(ctx context.Context, p rss.Post) (rss.Post, error) {
	post, err := c.rss.CreatePost(ctx, p)
	if err != nil {
		return rss.Post{}, fmt.Errorf("create: %w", err)
	}
	return post, err
}
