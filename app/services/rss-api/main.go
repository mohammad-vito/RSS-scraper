package main

import (
	"RssReader/app/services/rss-api/handlers"
	"RssReader/bussiness/data/store/rss"
	"RssReader/bussiness/sys/auth"
	"RssReader/bussiness/sys/config"
	"RssReader/foundation/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		config.DB.Host,
		config.DB.User,
		config.DB.Password,
		config.DB.Dbname,
		config.DB.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&rss.Feed{}, &rss.Post{})
	if err != nil {
		fmt.Println(err)
		return
	}
	r := gin.Default()

	log, err := logger.New("web")

	if err != nil {
		fmt.Println(err)
		return
	}
	a, _ := auth.New("mysecret_test")

	cfg := handlers.APIMuxConfig{
		Shutdown: nil,
		Log:      log,
		Auth:     &a,
		DB:       db,
	}

	handlers.V1(r, &cfg)

	r.Run() // listen

	//var t bool
	//db.Raw("select true").Scan(&t)
	//fmt.Println(t)
	//db.AutoMigrate(&user.User{})

}
