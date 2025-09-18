package middleware

import (
	"backend/pkg/jwt"
	"backend/pkg/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtService jwt.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(c, http.StatusUnauthorized, "Missing token", []response.APIError{
				{Code: "MISSING_TOKEN", Detail: "Authorization header is required"},
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			response.Error(c, http.StatusUnauthorized, "Invalid token format", []response.APIError{
				{Code: "INVALID_TOKEN_FORMAT", Detail: "Expected format: Bearer <token>"},
			})
			c.Abort()
			return
		}

		token, claims, err := jwtService.ValidateToken(parts[1])
		if err != nil || !token.Valid {
			response.Error(c, http.StatusUnauthorized, "Invalid or expired token", []response.APIError{
				{Code: "INVALID_TOKEN", Detail: err.Error()},
			})
			c.Abort()
			return
		}

		c.Set("user", claims)
		c.Next()
	}
}
