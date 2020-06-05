package enums

import (
	"os"
)

// GetEnv ... get the server env
func GetEnv() string {
	env := os.Getenv("SERVER_ENV")
	if env == "" {
		env = "development"
	}
	return env
}

// GetPort ... get the app port
func GetPort() string {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "2369"
	}
	return port
}

// GetAPIPrefix ... get the api prefix
func GetAPIPrefix() string {
	prefix := os.Getenv("API_PREFIX")
	if prefix == "" {
		prefix = "/api"
	}
	return prefix
}
