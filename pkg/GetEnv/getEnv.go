package GetEnv

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func load() {
	// Get the GITHUB_USERNAME environment variable
	githubUsername, exists := os.LookupEnv("GITHUB_USERNAME")

	if exists {
		fmt.Println(githubUsername)
	}

	// Get the GITHUB_API_KEY environment variable
	githubAPIKey, exists := os.LookupEnv("GITHUB_API_KEY")

	if exists {
		fmt.Println(githubAPIKey)
	}
	return (githubUsername)
}
