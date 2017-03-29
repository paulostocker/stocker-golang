package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("src/paulo.stocker/templates/public/login.html"))
		if err := t.ExecuteTemplate(w, "login.html", nil); err != nil { /// SelectData()
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("templates/error.html"))
		if err := t.ExecuteTemplate(w, "error.html", nil); err != nil { /// SelectData()
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
