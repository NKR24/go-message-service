package db

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func connect() *sqlx.DB {
	dsn := "host=localhost port=5432 user=youruser password=yourpassword dbname=yourdb sslmode=disable"
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
