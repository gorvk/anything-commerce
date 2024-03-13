package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://supergk:1525@localhost/anything_commerce_db"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		// _, e := db.Query("CALL create_user($1, $2, $3, $4, $5)", "gstar1525@gmail.com", "gourav", "kolhatkar", "9309704464", "Pune, India")
		rows, e := db.Query("SELECT * from read_all_users()")
		if e != nil {
			panic(e)
		}
		defer rows.Close()
		for rows.Next() {
			type User struct {
				id           int
				first_name   string
				last_name    string
				email        string
				phone_number string
				user_address string
				is_seller    bool
			}

			user := User{}

			if err := rows.Scan(
				&user.id,
				&user.first_name,
				&user.last_name,
				&user.email,
				&user.phone_number,
				&user.user_address,
				&user.is_seller); err != nil {
				panic(err)
			}

			fmt.Println(user.first_name)
		}
	}
}
