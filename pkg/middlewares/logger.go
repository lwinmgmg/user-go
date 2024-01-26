package middlewares

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware(c *gin.Context) {
	// Start timer
	start := time.Now()
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery

	// Process request
	c.Next()

	now := time.Now()
	if raw != "" {
		path = path + "?" + raw
	}
	statusCode := c.Writer.Status()
	slog.Default().Info(fmt.Sprintf("%15v %7v | %v | %13v | %7v | %v", c.ClientIP(), c.Request.Method, statusCode, now.Sub(start), c.Writer.Size(), path))
}
