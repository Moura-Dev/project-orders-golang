package main

import (
	"base-project-api/db"
	"base-project-api/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	database, _ := db.StartDB()
	database.MustExec(db.Schema)

	defer database.Close()

	s := server.NewServer()

	s.Run()
}
