package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
)

type Post struct {
	ID    int
	name string
	email  string
	pass string
}

var db, err = sql.Open("mysql", "root:root@/go_course?charset=utf8")

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("templates/index.html"))
		if err := t.ExecuteTemplate(w, "index.html", SelectData() ); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func InsertData(name, email, pass string)  {
	stmt, err := db.Prepare("INSERT INTO user(name, email, pass) values(?,?,?)")
	CheckError(err)

	_, err = stmt.Exec(name, email, pass)
	CheckError(err)

	db.Close();
}

func SelectData() []Post{
	
	rows, err := db.Query("SELECT * FROM user");
	CheckError(err)
	
	items := []Post{}
	for rows.Next() {
		post := Post{}

		rows.Scan(&post.ID, &post.name, &post.email, &post.pass)
		//fmt.Println(id, name, email, pass)
		items = append(items, post)
	}
	db.Close();
}