package middleware

import (
	"github.com/auto-calling/gateway/config"
	"github.com/auto-calling/gateway/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TokenValidate() gin.HandlerFunc {
	return func(c *gin.Context) {
		uri := c.Request.RequestURI
		token := c.Request.Header.Get("Authorization")
		if uri != "/api/ping" && token != config.TOKEN {
			handler.RespondWithError(c, http.StatusUnauthorized, "Invalid API token")
			return
		} else {
			c.Next()
		}
	}
}
