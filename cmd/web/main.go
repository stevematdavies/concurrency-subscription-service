package main

const WEB_PORT = "8080"

func main() {
	// Connect to a database
	db := InitDb()
	db.Ping()
	// Create Sessions
	session := InitSession()
	// Create Channels
	// Create a Wait Group
	// Setup Application Config
	// Set up mail
	// Listeb for Web connections
}
