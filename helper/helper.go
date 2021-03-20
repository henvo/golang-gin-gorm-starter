package helper

import "os"

// GetEnv finds an env variable or the given fallback.
func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}

	return value
}
