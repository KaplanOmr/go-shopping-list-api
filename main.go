package main

import "fmt"

func main() {

	var ad AppDB

	ad.conDB()

	defer ad.DB.Close()

	u := UserStruct{
		Username: "omer",
		Name:     "omer kaplan",
		Email:    "kaplanomer@outlook.com",
		Password: "123",
	}

	if !ad.createUser(u) {
		fmt.Println("user create error")
	}

	r, e := ad.DB.Query("SELECT * FROM users")

	if e != nil {
		panic(e)
	}

	var usr []UserStruct

	for r.Next() {
		var ou UserStruct
		if err := r.Scan(
			&ou.ID,
			&ou.Name,
			&ou.Username,
			&ou.Email,
			&ou.Password,
		); err != nil {
			fmt.Println(err)
		}

		usr = append(usr, ou)
	}

	fmt.Println(usr)

	// if err := fasthttp.ListenAndServe(fmt.Sprintf(":%s", settings["server"]["port"]), rootHandlers); err != nil {
	// 	panic(err)
	// }
}
