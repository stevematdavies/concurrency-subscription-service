package data

import (
	"database/sql"
	"time"
)

const dbTimeout = time.Second * 3

var db *sql.DB

type Models struct {
	User User
	Plan Plan
}

func New(dbp *sql.DB) Models {
	db = dbp
	return Models{
		User: User{},
		Plan: Plan{},
	}
}