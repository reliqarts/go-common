package common

import (
	"fmt"
	"os"
)

// GetEnvOrDefault gets environment variable or returns default fallback provided as second argument.
func GetEnvOrDefault(key string, defaultValue string) string {
	val := os.Getenv(key)

	if val == "" {
		val = defaultValue
	}

	return val
}

// MustGetEnv gets an environment variable or results in a panic.
func MustGetEnv(key string) string {
	val := GetEnvOrDefault(key, "")
	if val == "" {
		panic(fmt.Sprintf("Env var \"%s\" must be provided and is missing!", key))
	}

	return val
}
