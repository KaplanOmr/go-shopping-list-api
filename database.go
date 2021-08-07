package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func conDB() (*sql.DB, error) {
	s, err := getSettings()
	if err != nil {
		return nil, err
	}

	ct := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		s["database"]["hostname"],
		s["database"]["port"],
		s["database"]["username"],
		s["database"]["password"],
		s["database"]["database"])

	db, err := sql.Open("postgres", ct)
	if err != nil {
		return nil, err
	}

	return db, nil

}
