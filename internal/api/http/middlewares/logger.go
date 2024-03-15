package middlewares

import (
	"github.com/gin-gonic/gin"
	"log/slog"
)

func SlogLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path

		c.Next()

		statusCode := c.Writer.Status()
		slog.Info("Http Request", "status", statusCode, "path", path)
	}
}
