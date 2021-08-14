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

func (adb *AppDB) conDB() {

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

	adb.DB = db
}

func (adb *AppDB) createUser(u UserStruct) (int, bool) {

	r, err := adb.DB.Query(
		"SELECT username, email FROM users WHERE username = $1 OR email = $2",
		u.Username,
		u.Email,
	)

	if err != nil {
		fmt.Println(err)
		return 1, false
	}

	if r.Next() {
		var usr UserStruct
		err = r.Scan(&usr.Username, &usr.Password)
		if err != nil {
			return 1, false
		}

		if usr.Username == u.Username {
			return 2, false
		} else {
			return 3, false
		}
	}

	h := sha1.New()
	h.Write([]byte(u.Password))

	u.Password = fmt.Sprintf("%x", h.Sum(nil))

	fmt.Println(u.Password)
	_, err = adb.DB.Exec(
		"INSERT INTO users(username,name,email,password) VALUES ($1,$2,$3,$4)",
		u.Username,
		u.Name,
		u.Email,
		u.Password,
	)

	if err != nil {
		fmt.Println(err)
		return 1, false
	}

	return 0, true
}

func (adb *AppDB) checkUser(u UserStruct) (UserStruct, bool) {

	h := sha1.New()
	h.Write([]byte(u.Password))

	u.Password = fmt.Sprintf("%x", h.Sum(nil))

	r, err := adb.DB.Query(
		"SELECT id, username, name, email FROM users WHERE username = $1 AND password = $2 LIMIT 1",
		u.Username,
		u.Password,
	)

	if err != nil {
		fmt.Println(err)
		return UserStruct{}, false
	}

	if r.Next() {
		err = r.Scan(&u.ID, &u.Username, &u.Name, &u.Password)

		if err != nil {
			fmt.Println(err)
			return UserStruct{}, false
		}

		return u, true
	}

	return UserStruct{}, false
}
