package server

import (
	"github.com/Electra-project/electra-auth/src/controllers"
	"github.com/gin-gonic/gin"
)

// Router binds the routes to the controllers.
func Router() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"name":    "Electra Auth API",
				"version": "1",
			})
		})

		userGroup := v1.Group("user")
		{
			user := new(controllers.UserController)
			userGroup.GET("/:purseHash", user.Get)

		}
	}

	return router
}
