package main

import (
	"html/template"
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
	"paulo.stocker/models"
)

var port int

func main() {

	port = 666

	fmt.Printf("Initializing...\n")

	r := mux.NewRouter()
	r.HandleFunc("/", HandlerHome)
	r.HandleFunc("/error", HandlerError)

	fmt.Printf("Server listening on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}

func HandlerHome(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("src/paulo.stocker/views/public/login.html"))
	p := models.Product{Id: 1, Name: "Produto1", Description: "produto1 desc", Slug: "123"}
	fmt.Println(p.Name)

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
