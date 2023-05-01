package main

import (
	"RssReader/bussiness/core/crawler"
	"RssReader/bussiness/data/store/rss"
	"RssReader/bussiness/sys/message_broker"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	mb := message_broker.New()
	dsn := "host=localhost user=postgres password=postgres dbname=test port=5431 sslmode=disable TimeZone=Asia/Tehran"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	Fs := rss.NewStore(db)

	c := crawler.Config{Fs: &Fs, MB: mb}
	c.Run(2)

}
