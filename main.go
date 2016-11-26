package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"os"
	"api/controllers"
	"ws/watcher"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	hub := watcher.Run()

	router.HandleFunc("/api/suggestions", controllers.Suggestions).Methods("GET")
	router.HandleFunc("/ws",func(w http.ResponseWriter, r *http.Request) {
		watcher.Handle(hub, w, r)
	}).Methods("GET")

	http.Handle("/", router)

	log.Println("Listening at :" + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), nil))
}