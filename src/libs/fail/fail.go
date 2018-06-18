package fail

import (
	"net/http"
	"strings"

	"github.com/Electra-project/electra-auth/src/helpers"
	"github.com/gin-gonic/gin"
)

const (
	// MissingParameter error represents a missing parameter in the request query.
	MissingParameter = iota
	// MissingProperty error represents a missing property in the request body.
	MissingProperty
	// NotFound error represents an unfound entity in the database.
	NotFound
	// WrongPropertyValue error represents an unexpected property value.
	WrongPropertyValue
)

// Answer sends a JSON response error.
func Answer(c *gin.Context, err error) {
	errMessage := err.Error()
	helpers.LogErr(errMessage)

	switch true {

	case strings.Contains(errMessage, "json: cannot unmarshal"):
		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			gin.H{"message": "Wrong body properties types."},
		)

	default:
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"message": "Internal server error."},
		)
	}
}

// AnswerCustom sends a JSON response error related to a constant-defined error.
func AnswerCustom(c *gin.Context, errorIndex uint8, target string) {
	switch errorIndex {

	case MissingParameter:
		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			gin.H{"message": "Missing query parameter: " + target + "."},
		)

	case MissingProperty:
		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			gin.H{"message": "Missing body property: " + target + "."},
		)

	case NotFound:
		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			gin.H{"message": target + " not found."},
		)

	case WrongPropertyValue:
		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			gin.H{"message": "Wrong body property value for: " + target + "."},
		)

	default:
		helpers.LogErr(`Error: Couldn't handle this custom error.`)

		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"message": "Internal server error."},
		)
	}
}
