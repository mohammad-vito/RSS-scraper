package feedspb

import (
	"RssReader/bussiness/data/store/rss"
	"database/sql"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func (x *Feed) UnMarshall() rss.Feed {
	des := sql.NullString{
		String: "",
		Valid:  false,
	}
	if x.Description != nil {
		des.String = *x.Description
		des.Valid = true
	}

	lpt := sql.NullTime{
		Time:  time.Time{},
		Valid: false,
	}
	if x.LastPublishedAt != nil {
		lpt.Time = x.LastPublishedAt.AsTime()
		lpt.Valid = true
	}
	return rss.Feed{
		Link:            x.Link,
		Title:           x.Title,
		Description:     des,
		LastPublishedAt: lpt,
	}
}

func (x *Feed) Marshall(feed rss.Feed) {
	x.Id = int64(feed.ID)
	x.Link = feed.Link
	x.Title = feed.Title
	if feed.Description.Valid {
		x.Description = &feed.Description.String
	}
	if feed.LastPublishedAt.Valid {
		tp := timestamppb.New(feed.LastPublishedAt.Time)
		x.LastPublishedAt = tp
	}
}
