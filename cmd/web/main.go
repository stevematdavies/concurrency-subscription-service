package main

import (
	"log"
	"os"
	"subscriptionss/data"
	"sync"
)

const WEB_PORT = "8080"

func main() {
	db := InitDb()
	app := Config{
		Session: InitSession(),
		Db:db,
		Wg: &sync.WaitGroup{},
		InfoLog: log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		ErrorLog: log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		Models: data.New(db),
	}

	go app.ListenForShutdown()

	app.Serve()
}
