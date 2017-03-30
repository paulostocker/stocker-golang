package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", HandlerHome)
	r.HandleFunc("/error", HandlerError)
	http.ListenAndServe(":8080", r)
}

func HandlerHome(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("src/paulo.stocker/views/public/login.html"))
	if err := t.ExecuteTemplate(w, "login.html", nil); err != nil { /// SelectData()
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HandlerError(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("src/paulo.stocker/views/error.html"))
	if err := t.ExecuteTemplate(w, "error.html", nil); err != nil { /// SelectData()
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
