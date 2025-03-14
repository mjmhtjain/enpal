package main

import (
	"os"

	"github.com/mjmhtjain/enpal/src/router"
)

func main() {
	// Setup router with all routes configured
	router := router.Setup()

	// Start the server on port
	router.Run(":" + os.Getenv("PORT"))
}
