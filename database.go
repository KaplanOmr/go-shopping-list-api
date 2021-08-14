package main

import (
	"crypto/sha1"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type AppDB struct {
	DB *sql.DB
}

func (ad *AppDB) conDB() {

	ct := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		settings["database"]["hostname"],
		settings["database"]["port"],
		settings["database"]["username"],
		settings["database"]["password"],
		settings["database"]["database"])

	db, err := sql.Open("postgres", ct)
	if err != nil {
		fmt.Println(err)
	}

	ad.DB = db
}

func (ad *AppDB) createUser(u UserStruct) bool {

	h := sha1.New()
	h.Write([]byte(u.Password))

	u.Password = fmt.Sprintf("%x", h.Sum(nil))

	fmt.Println(u.Password)
	_, e := ad.DB.Exec(
		"INSERT INTO users(username,name,email,password) VALUES ($1,$2,$3,$4)",
		u.Username,
		u.Name,
		u.Email,
		u.Password,
	)

	if e != nil {
		fmt.Println(e)
		return false
	}

	return true
}
