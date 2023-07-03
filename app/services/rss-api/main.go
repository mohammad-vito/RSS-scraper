package main

import (
	"RssReader/app/services/rss-api/handlers"
	"RssReader/bussiness/sys/auth"
	"RssReader/bussiness/sys/config"
	"RssReader/foundation/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func run(log *zap.SugaredLogger) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		config.DB.Host,
		config.DB.User,
		config.DB.Password,
		config.DB.Dbname,
		config.DB.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	r := gin.Default()
	a, _ := auth.New("mysecret_test")

	cfg := handlers.APIMuxConfig{
		Shutdown: nil,
		Log:      log,
		Auth:     &a,
		DB:       db,
	}
	handlers.V1(r, &cfg)

	err = r.Run()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	log, err := logger.New("web")
	if err != nil {
		fmt.Printf("creating new logger: %w", err)
		os.Exit(1)
	}

	if err = run(log); err != nil {
		log.Errorw("startup", "ERROR", err)
		log.Sync()
		os.Exit(1)
	}
}
