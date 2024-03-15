package middlewares

import (
	"fmt"
	"github.com/BobrePatre/ProjectTemplate/internal/constants"
	"github.com/BobrePatre/ProjectTemplate/internal/utils/helpers"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strings"
)

const (
	op = "CORS Middleware"
)

// todo: Move to config
const (
	AllowOrigin     = "*" // more specific "localhost:3000, google.com"
	AllowCredential = "true"
	AllowHeader     = "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, User-Agent, Accept, Connection, Postman-Token" // separate with ", "
	AllowMethods    = "GET, PUT, DELETE, PATCH"
	MaxAge          = "43200" // for 12 hour
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", AllowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", AllowCredential)
		c.Writer.Header().Set("Access-Control-Allow-Headers", AllowHeader)
		c.Writer.Header().Set("Access-Control-Allow-Methods", AllowMethods)
		c.Writer.Header().Set("Access-Control-Max-Age", MaxAge)

		if !helpers.IsArrayContains(strings.Split(AllowMethods, ", "), c.Request.Method) {
			slog.Info(fmt.Sprintf("method %s is not allowed", c.Request.Method), slog.String(constants.LoggerCategoryMiddlewares, op), c.Request.Method)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden with CORS policy"})
			return
		}

		for key, value := range c.Request.Header {
			if !helpers.IsArrayContains(strings.Split(AllowHeader, ", "), key) {
				slog.Info(fmt.Sprintf("header %s is not allowed", key), slog.String(constants.LoggerCategoryMiddlewares, op), key, value)
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden with CORS policy"})
				return
			}
		}

		if AllowOrigin != "*" {
			if !helpers.IsArrayContains(strings.Split(AllowOrigin, ", "), c.Request.Host) {
				slog.Info(fmt.Sprintf("host '%s' is not part of '%v'\n", c.Request.Host, AllowOrigin), slog.String(constants.LoggerCategoryMiddlewares, op))
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden with CORS policy"})
				return
			}
		}

		c.Next()
	}
}
