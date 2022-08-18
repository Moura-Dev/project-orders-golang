package main

import (
	database "base-project-api/db"
	"base-project-api/server"
)

func main() {
	database.StartDB()

	s := server.NewServer()

	s.Run()
}
