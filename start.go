package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fp := path.Join("templates", "index.html")
	tmpl, _ := template.ParseFiles(fp)

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func InitRoutes() {
	r := mux.NewRouter()
	dir := "static"
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	r.HandleFunc("/", RootHandler)
	r.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	log.Print("Wassup")
	InitRoutes()
}
