package middlewares

import (
	"github.com/gin-gonic/gin"
)

// Cors sets the CORS headers.
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
	}
}
