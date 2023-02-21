package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const dbTimeOut = time.Second * 3

var db *sql.DB

func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		user: User{},
	}
}

type Models struct {
	user User
}

type User struct {
	Id        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Password  string    `json:"password"`
}
