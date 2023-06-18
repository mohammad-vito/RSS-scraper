package main

import (
	"RssReader/app/services/grpc/feedspb"
	"RssReader/bussiness/core/rss"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	feedspb.UnimplementedFeedServiceServer
	c *rss.Core
}

func (s *server) GetFeed(ctx context.Context, req *feedspb.GetFeedRequest) (*feedspb.GetFeedResponse, error) {
	feedId := int(req.GetId())
	feed, err := s.c.QueryByID(ctx, feedId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("feed %d", feed))
	}
	ts := timestamppb.New(feed.LastPublishedAt.Time)
	return &feedspb.GetFeedResponse{Feed: &feedspb.Feed{
		Id:              int64(feed.ID),
		Link:            feed.Link,
		Description:     &feed.Description.String,
		LastPublishedAt: ts,
	}}, nil
}

func (s *server) GetFeedList(ctx context.Context, req *feedspb.GetFeedListRequest) (*feedspb.GetFeedListResponse, error) {
	limit := int(req.GetLimit())
	offset := int(req.GetOffset())
	feeds, err := s.c.Query(ctx, offset, limit)
	if err != nil {
		// do sth
	}
	fs := make([]*feedspb.Feed, 0, len(feeds))
	for _, f := range feeds {
		fd := feedspb.Feed{}
		fd.Marshall(f)
		fs = append(fs, &fd)
	}

	resp := feedspb.GetFeedListResponse{
		Feeds: fs,
	}
	return &resp, nil

}
