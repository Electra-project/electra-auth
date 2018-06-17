package controllers

import (
	"github.com/Electra-project/electra-auth/src/models"
	"github.com/gin-gonic/gin"
)

// StatusController describes a status controller.
type StatusController struct{}

var statusModel = new(models.Status)

// Get a user public data.
func (s StatusController) Get(c *gin.Context) {
	status, err := statusModel.Get()
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error."})
		c.Abort()

		return
	}

	c.JSON(200, status)

	return
}
