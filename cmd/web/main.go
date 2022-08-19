package main

import (
	"log"
	"os"
	"sync"
)

const WEB_PORT = "8080"

func main() {
	app := Config{
		Session: InitSession(),
		Db: InitDb(),
		Wg: &sync.WaitGroup{},
		InfoLog: log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		ErrorLog: log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}

	go app.ListenForShutdown()

	app.Serve()
}
