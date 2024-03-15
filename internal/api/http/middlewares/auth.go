package middlewares

import (
	"errors"
	"github.com/BobrePatre/ProjectTemplate/internal/providers/web_auth_provider"
	authErrors "github.com/BobrePatre/ProjectTemplate/internal/providers/web_auth_provider/model"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strings"
)

func AuthMiddleware(provider web_auth_provider.WebAuthProvider) func(roles ...string) gin.HandlerFunc {
	return func(roles ...string) gin.HandlerFunc {
		return func(ctx *gin.Context) {
			authHeader := ctx.GetHeader("Authorization")

			if authHeader == "" {
				ctx.JSON(http.StatusUnauthorized, gin.H{"message": "No token provided"})
				ctx.Abort()
				return
			}

			headerArr := strings.Split(authHeader, " ")
			if len(headerArr) != 2 || headerArr[0] != "Bearer" {
				slog.Error("Invalid token format")
				ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token format, use Bearer {token}"})
				ctx.Abort()
				return
			}

			userDetails, err := provider.Authorize(ctx, headerArr[1], roles)
			if err != nil {
				switch {
				case errors.Is(err, authErrors.AccessDeniedError):
					ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Access denied"})
				default:
					ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
				}
				return
			}

			ctx.Set(web_auth_provider.UserDetailsKey, userDetails)
			ctx.Next()

		}
	}
}
