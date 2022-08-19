package main

import (
	"database/sql"
	"log"
	"subscriptionss/data"
	"sync"

	"github.com/alexedwards/scs/v2"
)

type Config struct {
	Session *scs.SessionManager
	Db *sql.DB
	InfoLog *log.Logger
	ErrorLog *log.Logger
	Wg * sync.WaitGroup
	Models data.Models
}