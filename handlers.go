package main

import (
	"fmt"
	"log"
)

func (app *App) simpleHandler(data string) {
	err := app.send(string(app.MenuItem.Type), fmt.Sprintf(app.MenuItem.Route, data), map[string]string{})
	if err != nil {
		log.Printf("Recieve an error: %s", err.Error())
	}
}