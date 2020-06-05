package db

import (
	"gousers/enums"

	"github.com/go-pg/pg"
)

// GetPgConnection ... intialize the DB
func GetPgConnection() *pg.DB {
	options := enums.GetPgOptions()
	db := pg.Connect(&options)
	return db
}
