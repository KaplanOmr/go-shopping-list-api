package main

import (
	"crypto/sha1"
	"database/sql"
	"errors"
	"fmt"
	"time"

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
		err = r.Scan(&u.ID, &u.Username, &u.Name, &u.Email)

		if err != nil {
			fmt.Println(err)
			return UserStruct{}, false
		}

		return u, true
	}

	return UserStruct{}, false
}

func (adb *AppDB) listCreate(u UserStruct, l ListStruct) (ListStruct, bool) {

	t := time.Now().Unix()

	var lid int

	err := adb.DB.QueryRow(
		"INSERT INTO lists(title,user_id,total_cost,status,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id;",
		l.Title,
		u.ID,
		l.TotalCost,
		l.Status,
		t,
		t,
	).Scan(&lid)

	if err != nil {
		fmt.Println(err)
		return ListStruct{}, false
	}

	l.ID = lid
	l.CreatedAt = int(t)

	return l, true
}

func (adb *AppDB) listGetAll(u UserStruct, desc, limit, offset string) ([]ListStruct, bool) {

	var lists []ListStruct

	if desc == "first" {
		desc = "ASC"
	} else {
		desc = "DESC"
	}

	sql := fmt.Sprintf(
		"SELECT id,title,total_cost,status,created_at,updated_at FROM lists WHERE deleted_at IS NULL AND user_id = $1 ORDER BY created_at %s LIMIT $2 OFFSET $3",
		desc,
	)

	result, err := adb.DB.Query(
		sql,
		u.ID,
		limit,
		offset,
	)

	if err != nil {
		fmt.Println(err)
		return []ListStruct{}, false
	}

	for result.Next() {
		var list ListStruct
		err = result.Scan(
			&list.ID,
			&list.Title,
			&list.TotalCost,
			&list.Status,
			&list.CreatedAt,
			&list.UpdatedAt,
		)

		if err != nil {
			fmt.Println(err)
			return []ListStruct{}, false
		}
		lists = append(lists, list)
	}

	return lists, true
}

func (adb *AppDB) listGet(u UserStruct, id int) (ListStruct, bool) {

	var l ListStruct

	err := adb.DB.QueryRow(
		"SELECT id, title, total_cost, status, created_at, updated_at FROM lists WHERE deleted_at IS NULL AND user_id = $1 AND id = $2 LIMIT 1",
		u.ID,
		id,
	).Scan(
		&l.ID,
		&l.Title,
		&l.TotalCost,
		&l.Status,
		&l.CreatedAt,
		&l.UpdatedAt,
	)

	if err != nil {
		fmt.Println(err)
		return ListStruct{}, false
	}

	return l, true
}

func (adb *AppDB) listUpdate(u UserStruct, title string, status, id int) bool {

	t := time.Now().Unix()

	up, err := adb.DB.Exec(
		"UPDATE lists SET title = $1, status = $2, updated_at = $3 WHERE id = $4 AND user_id = $5",
		title,
		status,
		t,
		id,
		u.ID,
	)

	if err != nil {
		fmt.Println(err)
		return false
	}

	count, err := up.RowsAffected()

	if err != nil {
		fmt.Println(err)
		return false
	}

	if count == 0 {
		fmt.Println(errors.New("ZERO AFFECT"))
		return false
	}

	return true
}

func (adb *AppDB) listDelete(u UserStruct, id int) bool {
	t := time.Now().Unix()

	del, err := adb.DB.Exec(
		"UPDATE lists SET deleted_at = $1 WHERE deleted_at IS NULL AND user_id = $2 AND id = $3",
		t,
		u.ID,
		id,
	)

	if err != nil {
		fmt.Println(err)
		return false
	}

	count, err := del.RowsAffected()

	if count == 0 {
		fmt.Println(errors.New("ZERO AFFECT"))
		return false
	}

	return true
}
