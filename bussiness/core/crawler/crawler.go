package crawler

import (
	"RssReader/bussiness/data/store/rss"
	"RssReader/bussiness/sys/message_broker"
	"context"
	"github.com/mmcdole/gofeed"
	"log"
	"sync"
)

const (
	topic          = "read-url"
	broker1Address = "localhost:9092"
)

type Config struct {
	Fs *rss.Store
	MB MessageBroker
}

type MessageBroker interface {
	SubmitMessages(ctx context.Context, TopicName string, key []byte, DTO ...interface{}) error
}

func (c Config) crawl(f rss.Feed, ctx context.Context) error {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(f.Link)
	if err != nil {
		return err
	}
	var posts []rss.Post
	for _, v := range feed.Items {
		if v.PublishedParsed.Before(f.LastPublishedAt.Time) {
			continue
		}
		post := rss.NewPost(v.Link, v.Title, "", v.Content, f.ID, *v.PublishedParsed)
		if err != nil {
			log.Fatalf("log failed %v", err)
		}

		posts = append(posts, post)

	}
	err = c.MB.SubmitMessages(ctx, message_broker.CrawlingTopic, []byte(feed.Title), posts)
	if err != nil {
		return err
	}
	return nil
}

func (c Config) Run(NumberOfWorkers int) error {
	ctx := context.Background()
	feeds := c.Fs.GetFeedsWithLastPublishedTime()
	ch := make(chan struct{}, NumberOfWorkers)
	var wg sync.WaitGroup
	for _, f := range feeds {
		wg.Add(1)
		go func(f rss.Feed) {
			ch <- struct{}{}
			log.Printf("crawling feed: %v %v", f.ID, f.Title)
			err := c.crawl(f, ctx)
			if err != nil {
				log.Printf("crawling feed failed: %v %v", f.ID, f.Title)
			} else {
				log.Printf("crawling feed success: %v %v", f.ID, f.Title)
			}
			<-ch
			wg.Done()
		}(f)
	}
	wg.Wait()
	close(ch)
	return nil
}
