package middleware

import (
	"github.com/gin-gonic/gin"
	"jwtTest/jwt"
	"net/http"
)

// Check TODO
func Check(context *gin.Context) {
	jwt.Secret = "123456"
	if token, ok := context.Request.Header["Token"]; ok {
		err := jwt.Verify(token[0])
		if err != nil {
			context.AbortWithStatusJSON(http.StatusForbidden, "403 Forbidden")
			return
		} else {
			payload, err := jwt.GetPayload(token[0])
			if err != nil {
				context.AbortWithStatusJSON(http.StatusInternalServerError, err)
			}
			context.Set("Username", payload.Audience)
			context.Next()
		}
	} else {
		context.AbortWithStatusJSON(http.StatusForbidden, "403 Forbidden")
	}
}
