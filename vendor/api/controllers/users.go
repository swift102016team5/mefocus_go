package controllers

import (
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType string `json:"token_type"`
}

func Authorization()string{
	url := "https://mefocus.auth0.com/oauth/token"

	payload := strings.NewReader("{\"client_id\":\"C3rfikNpiiljR3dFtLhS2NECpDtue0bK\",\"client_secret\":\"LnQvC92cKdiYeIiH7vFCIc3AN4Gt62UwvyF79U7TQUF8V1poc6xy9C8i7Ev4gdvN\",\"audience\":\"https://mefocus.auth0.com/api/v2/\",\"grant_type\":\"client_credentials\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var token Token
	err := json.Unmarshal(body, &token)
    if err != nil {
        println("Error parsing")
    }

	return token.TokenType + " " + token.AccessToken
}

func Users(w http.ResponseWriter, r *http.Request) {

	url := "https://mefocus.auth0.com/api/v2/users"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", Authorization())

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	w.Header().Set("Content-Type","application/vnd.api+json")
	w.Write(body)

}