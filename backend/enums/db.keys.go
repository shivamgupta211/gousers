package enums

import (
	"os"

	"github.com/go-pg/pg"
)

// GetPgOptions ... get the pg options
func GetPgOptions() pg.Options {
	dbURL := os.Getenv("DB_CONNECTION_URL")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_DATABASE")
	if dbURL == "" {
		dbURL = "localhost:5432"
	}
	if dbUser == "" {
		dbUser = "root"
	}
	if dbPassword == "" {
		dbPassword = "postgre"
	}
	if dbDatabase == "" {
		dbDatabase = "gousers"
	}
	return pg.Options{
		Addr:     dbURL,
		User:     dbUser,
		Password: dbPassword,
		Database: dbDatabase,
	}
}
