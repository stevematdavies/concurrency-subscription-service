package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func OpenDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func Connect() *sql.DB{
	attempts := 0
	dsn := os.Getenv("DSN")
	for {
		conn, err := OpenDb(dsn)
		if err != nil {
			log.Println("Postgres not ready....")
			attempts++
		} else {
			log.Println("Connected to Database")
			return conn
		}

		if attempts > 10 {
			return nil
		}

		log.Println("Backing off for one second...")
		time.Sleep(1 * time.Second)
		continue;
	}
}



func InitDb() *sql.DB{
	conn := Connect()
	if conn == nil {
		log.Panic("Cannot connect to database!")
	}
}