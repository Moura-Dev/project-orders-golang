package main

import (
	"github.com/joho/godotenv"

	"github.com/moura-dev/project-orders-golang/db"
	"github.com/moura-dev/project-orders-golang/server"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	database, _ := db.StartDB()

	defer database.Close()

	s := server.NewServer()

	s.Run()
}
