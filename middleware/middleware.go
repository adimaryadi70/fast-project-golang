package middleware

import (
	"fast-project-golang/tools"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err        := tools.TokenValid(c)
		if err != nil {
			//c.String(http.StatusUnauthorized, "Unauthorized")
			tools.ResAll(c,"","01","Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}