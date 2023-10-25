package middlewares

import (
	"gorm-practice/models/dto"
	"gorm-practice/utils/exception"
	"gorm-practice/utils/security"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// type authHeader struct {
// 	AuthorizationHeader string `header:"Authorization"`
// }

// Middleware sederhana untuk otentikasi
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")

		if authorization == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Status:  exception.StatusUnauthorized,
				Message: exception.ErrTokenNotProvided.Error(),
			})
			return
		}

		token := strings.Split(authorization, " ")[1]

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Status:  exception.StatusUnauthorized,
				Message: exception.ErrTokenRequired.Error(),
			})
			return
		}

		if _, blackListed := security.BlackListedTokens[token]; blackListed {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Status:  exception.StatusUnauthorized,
				Message: exception.ErrInvalidToken.Error(),
			})
			return
		}

		claims, err := security.VerifyAccessToken(token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Status:  exception.StatusUnauthorized,
				Message: err.Error(),
			})
			return
		}

		c.Set("username", claims["Username"])

		c.Next()
	}
}
