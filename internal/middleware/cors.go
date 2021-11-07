package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("AllowOrigins", "POST, GET, OPTIONS, PUT, DELETE")
	c.Writer.Header().Set("AllowHeaders", "*")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "Accept, Content-Type, Content-Length, Accept-Encoding")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	c.Next()
}

