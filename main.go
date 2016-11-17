package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"os"

	"api/controllers"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", controllers.Suggestions).Methods("GET")

	http.Handle("/", r)

	log.Println("Listening at :" + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), nil))
}