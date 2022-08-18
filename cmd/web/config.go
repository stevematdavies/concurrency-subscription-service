package main

import (
	"database/sql"
	"log"
	"sync"

	"github.com/alexedwards/scs/v2"
)

type Config struct {
	Session *scs.SessionManager
	Db *sql.DB
	InfoLog *log.Logger
	ErrorLog *log.Logger
	Wg * sync.WaitGroup
}