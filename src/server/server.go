package server

import (
	"os"

	"github.com/joho/godotenv"
)

// Start the server.
func Start() {
	if os.Getenv("GO_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}

	router := Router()
	router.Run(":" + getPort())
}

func getPort() string {
	envPort := os.Getenv("PORT")

	if envPort == "" {
		return "8080"
	}

	return envPort
}
