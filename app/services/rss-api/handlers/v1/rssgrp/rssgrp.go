package rssgrp

import "C"
import (
	rssCore "RssReader/bussiness/core/rss"
	"RssReader/bussiness/data/store/rss"
	"RssReader/bussiness/sys/validate"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func (h Handler) Query(c *gin.Context) {
	strOffset := c.Param("offset")
	Offset, err := strconv.Atoi(strOffset)
	if err != nil {
		Offset = 0
	}

	strLimit := c.Param("limit")
	Limit, err := strconv.Atoi(strLimit)
	if err != nil {
		Limit = 20
	}

	feeds, err := h.Rss.Query(c, Offset, Limit)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, feeds)
}

func (h Handler) QueryByID(c *gin.Context) {
	strID := c.Param("id")
	if strID == "" {
		c.Abort()
	}
	ID, err := strconv.Atoi(strID)
	if err != nil {
		c.Abort()
		return
	}
	feed, err := h.Rss.QueryByID(c, ID)
	if err != nil {
		switch er := validate.Cause(err); {
		case er == gorm.ErrRecordNotFound:
			c.Error(validate.NewRequestError(er, http.StatusNotFound))
			return
		}
	}
	c.JSON(http.StatusOK, feed)
}
