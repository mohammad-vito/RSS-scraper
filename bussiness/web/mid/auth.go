package mid

import (
	"RssReader/bussiness/sys/auth"
	"RssReader/bussiness/sys/validate"
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
			err = validate.NewRequestError(err, http.StatusBadRequest)
			context.Error(err)
			context.Abort()
		}
		claims, err := a.ValidateToken(parts[1])
		if err != nil {
			err := errors.New("unauthorized")
			err = validate.NewRequestError(err, http.StatusUnauthorized)
			context.Error(err)
			context.Abort()
		}
		auth.SetClaims(context, claims)
		context.Next()
	}
}
