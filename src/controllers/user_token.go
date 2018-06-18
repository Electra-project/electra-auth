package controllers

import (
	"github.com/gin-gonic/gin"
)

// UserTokenController class.
type UserTokenController struct{}

// Get a user token challenge from a Purse Account address hash.
func (u UserTokenController) Get(c *gin.Context) {
	purseHash := c.Param("purseHash")

	if !isPurseHashValid(purseHash) {
		c.JSON(400, gin.H{"message": "Invalid Purse Account address hash."})
		c.Abort()

		return
	}

	user, err := userTokenModel.GetByPurseHash(c.Param("purseHash"))
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error."})
		c.Abort()

		return
	}

	c.JSON(200, gin.H{"user": user})

	return
}

func isPurseHashValid(purseHash string) bool {
	return len(purseHash) == 34 && string([]rune(purseHash)[0]) == "E"
}
