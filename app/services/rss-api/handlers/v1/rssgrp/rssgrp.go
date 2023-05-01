package rssgrp

import (
	rssCore "RssReader/bussiness/core/rss"
	"RssReader/bussiness/data/store/rss"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct {
	Rss rssCore.Core
}

func (h Handler) Create(c *gin.Context) {
	var feed rss.Feed
	err := c.BindJSON(&feed)
	if err != nil {
		c.Error(err)
		return
	}
	feed, err = h.Rss.CreateFeed(c, feed)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, feed)
}

func (h Handler) QueryByID(c *gin.Context) {
	strID := c.Param("id")
	if strID == "" {
		c.Abort()
	}
	ID, err := strconv.Atoi(strID)
	if err != nil {
		c.Abort()
	}
	feed, err := h.Rss.QueryByID(c, ID)
	if err != nil {
		c.Abort()
	}
	c.JSON(http.StatusOK, feed)
}
