package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"fullStack/app/helpers"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Otorisasi")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Tidak sah"})
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Format token tidak valid"})
			c.Abort()
			return
		}
		tokenString := tokenParts[1]

		userID, err := helpers.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			c.Abort()
			return
		}

		c.Set("userID", userID)

		c.Next()
	}
}
