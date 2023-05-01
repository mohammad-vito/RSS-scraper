package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Respond(c *gin.Context, data interface{}, statusCode int) {
	if statusCode == http.StatusNoContent {
		c.Done()
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(statusCode, data)
}
