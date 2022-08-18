package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gomodule/redigo/redis"
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
		} else {
			log.Println("Connected to Database")
			return conn
		}

		if attempts > 10 {
			return nil
		}

		log.Println("Backing off for one second...")
		time.Sleep(1 * time.Second)
		attempts++
		continue;
	}
}

func InitDb() *sql.DB{
	conn := Connect()
	if conn == nil {
		log.Panic("Cannot connect to database!")
	}
	return conn
}

func InitSession() *scs.SessionManager{
	s := scs.New()
	s.Store = redisstore.New(RedisConnect())
	s.Lifetime = 24 * time.Hour
	s.Cookie.Persist = true
	s.Cookie.SameSite = http.SameSiteLaxMode
	s.Cookie.Secure = true
	return s
}

func RedisConnect() *redis.Pool{
	pool := &redis.Pool{
		MaxIdle: 10,
		Dial: func()(redis.Conn, error) {
			return redis.Dial("tcp", os.Getenv("REDIS"))
		},
	}
	return pool
}

func (app *Config) Serve() {
	s := &http.Server{
		Addr: fmt.Sprintf(":%s", WEB_PORT),
		Handler: app.Routes(),
	}
	app.InfoLog.Println("Starting Web Server...")
	if err := s.ListenAndServe(); err != nil {
		log.Panic(err)
	}

}