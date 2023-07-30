package router

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		log.Printf(
			"%s - [%s] \"%s %s %s\" %d %s",
			c.ClientIP(),
			end.Format("02/Jan/2006:15:04:05 -0700"),
			c.Request.Method,
			c.Request.URL.Path,
			c.Request.Proto,
			c.Writer.Status(),
			latency.String(),
		)
	}
}
