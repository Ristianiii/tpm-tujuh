package middleware

import (
	"net/http"
	"strings"
	"tpm-tujuh/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func BearerAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerAuth := ctx.GetHeader("Authorization")
		splitToken := strings.Split(headerAuth, " ")
		if len(splitToken) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]any{
				"message": "invalid authorization header",
			})
			return
		}

		// check basic
		if splitToken[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]any{
				"message": "invalid authorization method",
			})
			return
		}

		// validate jwt
		valid, token := service.ValidateUserJWT(splitToken[1])
		if !valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]any{
				"message": "malformed token",
			})
			return
		}

		// get claim
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]any{
				"message": "failed to parse claims",
			})
			return
		}

		// add value in context
		userID, ok := claims["uid"].(string)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "Failed to get user_id from JWT claims",
			})
			return
		}
		ctx.Set("user_id", userID)
		ctx.Next()
	}
}
