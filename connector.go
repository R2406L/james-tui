package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func (app *App) send(method string, route string, args map[string]string) (err error) {
	url := app.Url + route
	log.Printf("Request URI %s", url)
	
	postBody, err := json.Marshal(args)
	if err != nil {
		return
	}

	postBodyBytes := bytes.NewBuffer(postBody)

	resp, err := http.NewRequest(method, url, postBodyBytes)
	if err != nil {
		return
	}

	response, err := http.DefaultClient.Do(resp)
	if err != nil {
		return
	}

	app.Response = response
	return
}