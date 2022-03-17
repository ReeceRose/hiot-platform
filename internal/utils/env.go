package utils

import (
	"os"
)

//GetEnvironmentVariable attempts to get an environment variable.
//If no environment variable is set, it will return the fallback
func GetEnvironmentVariable(key string, fallback string) string {
	if variable := os.Getenv(key); variable != "" {
		return variable
	}
	return fallback
}
