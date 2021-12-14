package handler

import "github.com/gin-gonic/gin"

func RespondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"msg": message, "success": false})
}
