package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func conDB() (*sql.DB, error) {
	getSettings()

	ct := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		settings["database"]["hostname"],
		settings["database"]["port"],
		settings["database"]["username"],
		settings["database"]["password"],
		settings["database"]["database"])

	db, err := sql.Open("postgres", ct)
	if err != nil {
		return nil, err
	}

	return db, nil

}
