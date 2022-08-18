package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)



func StartDB() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	// migrations.RunMigrations(db)
	fmt.Println("Database connected", db)
}
