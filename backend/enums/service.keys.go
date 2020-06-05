package enums

import (
	"os"
)

// GetInventoryServiceAPI ... get the sentry link
func GetInventoryServiceAPI() string {
	api := os.Getenv("INVENTORY_SERVICE_API")
	if api == "" {
		api = "http://localhost:1369/api"
	}
	return api
}
