package main

import (
	"RssReader/app/services/grpc/feedspb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	c := feedspb.NewFeedServiceClient(conn)
	res, err := c.GetFeedList(context.Background(), &feedspb.GetFeedListRequest{
		Limit:  5,
		Offset: 0,
	})
	fmt.Println(res.Feeds)
}
