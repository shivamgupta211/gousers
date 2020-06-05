package usermodule

import (
	"time"
)

// User ...model for User
type User struct {
	ID        int       `pg:"id" json:"id"`
	FirstName string    `pg:"first_name" json:"firstName"`
	LastName  string    `pg:"last_name" json:"lastName"`
	Skills    string    `pg:"skills" json:"skills"`
	Interests string    `pg:"interests" json:"interests"`
	CreatedAt time.Time `pg:"created_at,default:now()" json:"created_at"`
	UpdatedAt time.Time `pg:"updated_at,default:now()" json:"updated_at"`
}
