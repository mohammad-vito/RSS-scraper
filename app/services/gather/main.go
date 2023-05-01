package main

import (
	"RssReader/bussiness/data/store/rss"
	"RssReader/bussiness/sys/message_broker"
	"context"
	"fmt"
	"sync"
)

func main() {

	mb := message_broker.New()
	var wg sync.WaitGroup
	//dsn := "host=localhost user=postgres password=postgres dbname=test port=5431 sslmode=disable TimeZone=Asia/Tehran"
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	panic(err)
	//}
	//core := rss_core.NewCore(db)
	//ctx := context.Background()
	var p []rss.Post
	ch := mb.Listen(context.Background(), message_broker.CrawlingTopic, &p, &wg)
	for v := range ch {
		fmt.Println(v)
		//post, _ := v.(*rss.Post)
		//pp, err := core.CreatePost(ctx, *post)
		//if err != nil {
		//	log.Println(err)
		//} else {
		//	log.Printf("post %v of feed %v added.", pp.ID, pp.FeedID)
		//}
	}
	wg.Wait()

}
