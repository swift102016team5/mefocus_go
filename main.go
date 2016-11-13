package main

import (
	"html/template"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Index request", r.RemoteAddr, r.URL)
	t, _ := template.ParseFiles("views/index.html")
	t.Execute(w, nil)
}

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", index).Methods("GET")

	http.Handle("/", r)

	log.Println("Listening at :" + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), nil))
}