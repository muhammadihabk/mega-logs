package utilities

import (
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	ErrorHandler(err, "Failed to read in the environment variables.")
}
