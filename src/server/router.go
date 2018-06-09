package server

import "github.com/gin-gonic/gin"

// Router binds the routes to the controllers.
func Router() *gin.Engine {
	router := gin.Default()

	router.Group("v1.0")
	{
		router.GET("/v", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"name":    "Electra Auth API",
				"version": "1.0",
			})
		})
	}

	return router
}
