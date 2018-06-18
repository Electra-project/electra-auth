package controllers

import (
	"net/http"

	"github.com/Electra-project/electra-auth/src/libs/fail"
	"github.com/gin-gonic/gin"
)

// UserController describes a user controller.
type UserController struct{}

type userPostBody struct {
	Signature string `json:"signature"`
}

// Get retrieves the authenticated user info.
func (u UserController) Get(c *gin.Context) {
	purseHash := getPurseHash(c)

	user, err := userModel.GetByPurseHash(purseHash)
	if err != nil {
		fail.Answer(c, err, "user")

		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})

	return
}

// Post creates a new database entry for the authenticated user.
func (u UserController) Post(c *gin.Context) {
	purseHash := getPurseHash(c)

	var reqBody *userPostBody
	err := c.BindJSON(&reqBody)
	if err != nil {
		fail.Answer(c, err, "user")

		return
	}

	user, err := userModel.Insert(purseHash)
	if err != nil {
		fail.Answer(c, err, "user")

		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": user})
}
