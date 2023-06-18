package main

import (
	"RssReader/app/services/grpc/feedspb"
	"RssReader/bussiness/core/rss"
	"RssReader/bussiness/sys/config"
	"RssReader/foundation/logger"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net"
)

func run(lgr *zap.SugaredLogger) error {
	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
		// Add any other option (check functions starting with logging.With).
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		config.DB.Host,
		config.DB.User,
		config.DB.Password,
		config.DB.Dbname,
		config.DB.Port)

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
	return nil
}

func main() {
	lgr, err := logger.New("grpc")
	if err != nil {
		// do
	}
	if err = run(lgr); err != nil {
		// do sth
	}
}
