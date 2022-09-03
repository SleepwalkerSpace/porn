package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		begin := time.Now()
		c.Next()
		logger.WithFields(logrus.Fields{
			"Status":  c.Writer.Status(),
			"Method":  c.Request.Method,
			"Path":    c.Request.URL.Path,
			"Elapsed": time.Since(begin).String(),
			"IP":      c.ClientIP(),
		}).Printf("%+v", c.Request.Body)
	}
}
