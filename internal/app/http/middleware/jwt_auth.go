package middleware

import (
	"net/http"
	"strings"

	"goframe/internal/core/config"
	"goframe/internal/core/jwt"

	"github.com/gin-gonic/gin"
)

func JWTAuth(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization header is missing",
			})
			return
		}

		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization header format must be Bearer {token}",
			})
			return
		}

		tokenStr := parts[1]
		claims, err := jwt.ParseAccessToken(cfg.JWTSecret, tokenStr)
		if err != nil {
			errMsg := err.Error()
			if strings.Contains(errMsg, "token is expired") {
				errMsg = "token has expired"
			} else if strings.Contains(errMsg, "signature is invalid") {
				errMsg = "token signature is invalid"
			} else {
				errMsg = "invalid token"
			}

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"error":   errMsg,
			})
			return
		}

		// Token valid
		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Next()
	}
}
