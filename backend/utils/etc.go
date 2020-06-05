package utils

import (
	"os"
	"time"
)

// GetNowUTC - get now as string
func GetNowUTC() time.Time {
	now := time.Now().UTC()
	return now
}

// ReadENV - Read env var
func ReadENV(key, defaultVal string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultVal
	}
	return value
}
