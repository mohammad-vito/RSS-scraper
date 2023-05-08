package main

import (
	"RssReader/bussiness/core/crawler"
	"RssReader/bussiness/data/store/rss"
	"RssReader/bussiness/sys/config"
	"RssReader/bussiness/sys/message_broker"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	mb := message_broker.New()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		config.DB.Host,
		config.DB.User,
		config.DB.Password,
		config.DB.Dbname,
		config.DB.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	Fs := rss.NewStore(db)

	c := crawler.Config{Fs: &Fs, MB: mb}
	c.Run(2)
}
