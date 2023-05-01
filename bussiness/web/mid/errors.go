package mid

import (
	"RssReader/bussiness/sys/validate"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func Errors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			var err validate.ErrorResponse
			var status int
			switch e := validate.Cause(c.Errors[0]).(type) {
			case validator.ValidationErrors:
				for _, v := range e {
					fe := validate.ErrorField{
						Field: v.Field(),
						Err:   v.Error(),
					}
					err.Fields = append(err.Fields, fe)
				}
				err.Error = "data validation error."
				status = http.StatusBadRequest
			case *validate.RequestError:
				err.Error = e.Error()
				status = e.Status

			default:
				err.Error = "internal server error."
				status = http.StatusInternalServerError
			}
			c.JSON(status, err)

		}

	}
}
