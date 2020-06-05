package enums

import (
	"os"
)

// GetSentry ... get the sentry link
func GetSentry() string {
	sentry := os.Getenv("SENTRY")
	if sentry == "" {
		sentry = "https://somedummy@sentry.io/Gousers"
	}
	return sentry
}
