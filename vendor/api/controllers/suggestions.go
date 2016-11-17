package controllers

import (
	"encoding/json"
	"net/http"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type Suggestion struct {
	ID        		bson.ObjectId 	`json:"id" bson:"_id"`
	Name      		string 			`json:"name"`
	Duration     	string 			`json:"duration"`
}

type SuggestionsResponse struct {
	Data []Suggestion `json:"data"`
}

func Suggestions(w http.ResponseWriter, r *http.Request) {
	session, err := mgo.Dial("mongodb://mefocus:mefocus@ds157247.mlab.com:57247/mefocus")
    if err != nil {
        panic(err)
    }

    defer session.Close()
    
    session.SetMode(mgo.Monotonic, true)

    collection := session.DB("mefocus").C("suggestions")
	result := []Suggestion{}

	err = collection.Find(nil).All(&result)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type","application/vnd.api+json")
	json.NewEncoder(w).Encode(SuggestionsResponse{result})
}