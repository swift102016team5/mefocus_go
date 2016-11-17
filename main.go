package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"os"

	"api/controllers"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/suggestions", controllers.Suggestions).Methods("GET")

	http.Handle("/", router)

	log.Println("Listening at :" + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), nil))
}