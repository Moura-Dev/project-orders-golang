package main

import (
	"base-project-api/database"
	"base-project-api/server"
)

func main() {
	database.StartDB()
	s := server.NewServer()

	s.Run()
}
