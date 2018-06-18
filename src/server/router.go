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

		status := new(controllers.StatusController)
		v1.GET("/status", status.Get)

		userGroup := v1.Group("user")
		{
			user := new(controllers.UserController)
			userGroup.GET("/:purseHash", user.Get)
			userGroup.POST("/:purseHash", user.Post)

			userToken := new(controllers.UserTokenController)
			userGroup.GET("/:purseHash/token", userToken.Get)
		}
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not Found"})
	})

	return router
}
