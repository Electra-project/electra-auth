package main

import (
	"os"

	"github.com/Electra-project/electra-auth/src/server"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("GO_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}

	server.Start()
}
