package controllers

import (
	"github.com/Electra-project/electra-auth/src/models"
	"github.com/gin-gonic/gin"
)

// UserController describes a user controller.
type UserController struct{}

var userModel = new(models.User)

// Get a user public data.
func (u UserController) Get(c *gin.Context) {
	if c.Param("purseHash") != "" {
		user, err := userModel.GetByPurseHash(c.Param("purseHash"))
		if err != nil {
			if err.Error() == "not found" {
				c.JSON(404, gin.H{"message": "User not found."})
			} else {
				c.JSON(500, gin.H{"message": "Internal Server Error."})
			}
			c.Abort()

			return
		}

		c.JSON(200, gin.H{"user": user})

		return
	}

	c.JSON(400, gin.H{"message": "Bad Request."})
	c.Abort()

	return
}
