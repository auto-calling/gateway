package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetDurationInMillSeconds(start time.Time) float64 {
	end := time.Now().UTC()
	duration := end.Sub(start)
	milliseconds := float64(duration) / float64(time.Millisecond)
	rounded := float64(int(milliseconds*100+.5)) / 100
	return rounded
}

// JSONLogMiddleware logs a gin HTTP request in JSON format, with some additional custom key/values
func JSONLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now().UTC()

		// Process Request
		c.Next()

		entry := log.WithFields(log.Fields{
			"src_ip":    c.Request.RemoteAddr,
			"duration":  GetDurationInMillSeconds(start),
			"method":    c.Request.Method,
			"path":      c.Request.RequestURI,
			"status":    c.Writer.Status(),
			"referrer":  c.Request.Referer(),
			"client_ip": c.Request.Header.Get("client-ip"),
			"protocol": c.Request.Proto,
			"user-agent": c.Request.UserAgent(),
		})

		if c.Writer.Status() >= 500 {
			entry.Error(c.Errors.String())
		} else {
			entry.Info("")
		}
	}
}
