package models

import (
	"database/sql"

	_ "github.com/mysql"
	"paulo.stocker/generic"
)

type User struct {
	ID    int
	name  string
	email string
	pass  string
}

var db, err = sql.Open("mysql", "root:root@/go_course?charset=utf8")

func InsertData(name, email, pass string) {
	stmt, err := db.Prepare("INSERT INTO user(name, email, pass) values(?,?,?)")
	generic.CheckError(err)

	_, err = stmt.Exec(name, email, pass)
	generic.CheckError(err)
	db.Close()
}

func SelectData() []User {

	rows, err := db.Query("SELECT * FROM user")
	generic.CheckError(err)
	items := []User{}
	for rows.Next() {
		user := User{}

		rows.Scan(&user.ID, &user.name, &user.email, &user.pass)
		//fmt.Println(id, name, email, pass)
		items = append(items, user)
	}
	db.Close()
	return items
}
