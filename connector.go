package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func (app *App) send(method string, route string, args map[string]string) (err error) {
	url := app.Url + route
	log.Printf("Request URI %s", url)
	
	ctx, cncl := context.WithTimeout(context.Background(), time.Second*3)
	defer cncl()

	postBody, err := json.Marshal(args)
	if err != nil {
		return
	}

	postBodyBytes := bytes.NewBuffer(postBody)

	resp, err := http.NewRequestWithContext(ctx, method, url, postBodyBytes)
	if err != nil {
		return
	}

	response, err := http.DefaultClient.Do(resp)
	if err != nil {
		return
	}
	
	app.Response = response
	app.setResponse()

	return
}