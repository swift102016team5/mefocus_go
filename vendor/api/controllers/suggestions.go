package controllers

import (
	"log"
	"net/http"
	"html/template"
)

func Suggestions(w http.ResponseWriter, r *http.Request) {
	log.Println("Index request", r.RemoteAddr, r.URL)
	t, _ := template.ParseFiles("views/index.html")
	t.Execute(w, nil)
}