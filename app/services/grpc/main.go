package main

import (
	"RssReader/app/services/grpc/feedspb"
	"RssReader/bussiness/core/rss"
	"RssReader/foundation/logger"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net"
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

func main() {
	lgr, err := logger.New("grpc")
	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
		// Add any other option (check functions starting with logging.With).
	}
	dsn := "host=localhost user=postgres password=postgres dbname=test port=5431 sslmode=disable TimeZone=Asia/Tehran"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	core := rss.NewCore(db)
	lis, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(InterceptorLogger(lgr), opts...),
		),
		grpc.ChainStreamInterceptor(
			logging.StreamServerInterceptor(InterceptorLogger(lgr), opts...),
		),
	)
	feedspb.RegisterFeedServiceServer(s, &server{c: &core})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
