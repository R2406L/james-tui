package main

import (
	"fmt"
	"log"
)

func (app *App) simpleHandler(args []string) {
	err := app.send(string(app.MenuItem.Type), fmt.Sprintf(app.MenuItem.Route, args[0]), map[string]string{})
	if err != nil {
		log.Printf("Recieve an error: %s", err.Error())
	}
}

func (app *App) inputEmailPasswordHandler(args []string) {
	log.Printf("Recieved args: %v", args)
}