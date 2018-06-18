package server

import (
	"github.com/Electra-project/electra-auth/src/controllers"
	"github.com/Electra-project/electra-auth/src/middlewares"
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

		statusController := new(controllers.StatusController)
		v1.GET("/status", statusController.Get)

		userGroup := v1.Group("user")
		{
			userTokenController := new(controllers.UserTokenController)
			userGroup.GET("/:purseHash/token", userTokenController.Get)
			userGroup.POST("/:purseHash/token", userTokenController.Post)
		}

		v1.Use(middlewares.IsUser())
		{
			userController := new(controllers.UserController)
			v1.GET("/user", userController.Get)
			v1.POST("/user", userController.Post)
		}
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not Found"})
	})

	return router
}
