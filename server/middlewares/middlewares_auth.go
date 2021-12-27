package middlewares

import (
	"pedy/services/auth"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const Bearer_schema = "Bearer "
		header := c.GetHeader("Authorization")

		if header == "" {
			c.AbortWithStatus(401)
		}

		token := header[len(Bearer_schema):]
		if !auth.NewJwtService().ValidateToken(token) {
			c.AbortWithStatus(401)
		}
	}
}
