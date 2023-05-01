package mid

import (
	"RssReader/bussiness/sys/auth"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Authenticate(a *auth.Auth) gin.HandlerFunc {
	return func(context *gin.Context) {
		authStr := context.Request.Header.Get("Authorization")
		parts := strings.Split(authStr, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			err := errors.New("expected authorization header format: bearer <token>")
			context.AbortWithError(http.StatusUnauthorized, err)
		}
		claims, err := a.ValidateToken(parts[1])
		if err != nil {
			context.AbortWithError(http.StatusUnauthorized, err)
		}
		auth.SetClaims(context, claims)
		context.Next()
	}
}
