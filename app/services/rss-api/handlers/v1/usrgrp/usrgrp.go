package usrgrp

import (
	userCore "RssReader/bussiness/core/user"
	"RssReader/bussiness/data/store/user"
	"RssReader/bussiness/sys/auth"
	"RssReader/foundation/web"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Handler struct {
	User userCore.Core
	Auth *auth.Auth
}

func (h Handler) Register(c *gin.Context) {
	var u user.NewUser
	if err := c.BindJSON(&u); err != nil {
		panic(err)
		return
	}
	usr, err := h.User.Create(c, u, time.Now())
	if err != nil {
		c.JSON(http.StatusInternalServerError, struct {
			Message string
		}{Message: "schema not accepted"})
		return
	}
	web.Respond(c, usr, http.StatusCreated)
}

func (h Handler) Login(c *gin.Context) {
	var data struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Message string
		}{Message: "schema not accepted"})
		return
	}
	claims, err := h.User.Authenticate(c, time.Now(), data.Email, data.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Message string
		}{Message: "invalid username or password"})
		return
	}
	token, err := h.Auth.GenerateToken(claims)
	if err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Message string
		}{Message: "invalid username or password"})
	}
	c.JSON(http.StatusBadRequest, struct {
		Token string `json:"access_token"`
	}{Token: token})
}
