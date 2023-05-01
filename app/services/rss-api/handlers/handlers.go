package handlers

import (
	"RssReader/app/services/rss-api/handlers/v1/rssgrp"
	"RssReader/app/services/rss-api/handlers/v1/usrgrp"
	rssCore "RssReader/bussiness/core/rss"
	userCore "RssReader/bussiness/core/user"
	"RssReader/bussiness/sys/auth"
	"RssReader/bussiness/web/mid"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
	Auth     *auth.Auth
	DB       *gorm.DB
}

func V1(r *gin.Engine, cfg *APIMuxConfig) {
	r.Use(mid.Errors())
	ugh := usrgrp.Handler{
		User: userCore.NewCore(cfg.Log, cfg.DB),
		Auth: cfg.Auth,
	}

	r.POST("/register", ugh.Register)
	r.POST("/login", ugh.Login)

	fgh := rssgrp.Handler{
		Rss: rssCore.NewCore(cfg.DB),
	}

	r.POST("/feeds", mid.Authenticate(cfg.Auth), fgh.Create)
	r.GET("/feeds", mid.Authenticate(cfg.Auth), fgh.Query)
	r.GET("/feeds/:id", mid.Authenticate(cfg.Auth), fgh.QueryByID)

}
