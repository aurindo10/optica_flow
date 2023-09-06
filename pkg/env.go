package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var loaded = false

// LoadEnvironment func to load the .env file
func LoadEnvironment() error {
	if loaded {
		return nil
	}
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}
	loaded = true
	return nil
}

// Config retrieves the value of the environment variable associated with the key.
func Config(key string) string {
	if !loaded {
		err := LoadEnvironment()
		if err != nil {
			fmt.Println(err)
			return ""
		}
	}
	return os.Getenv(key)
}
