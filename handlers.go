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

func (app *App) addUserHandler(args []string) {
	err := app.send(string(app.MenuItem.Type), fmt.Sprintf(app.MenuItem.Route, args[0]), map[string]string{
		"password": args[1],
	})
	if err != nil {
		log.Printf("Recieve an error: %s", err.Error())
	}
}

func (app *App) changePasswordHandler(args []string) {
	err := app.send(string(app.MenuItem.Type), fmt.Sprintf(app.MenuItem.Route, args[0]), map[string]string{
		"password": args[1],
	})
	if err != nil {
		log.Printf("Recieve an error: %s", err.Error())
	}
}

func (app *App) mailboxClearHandler(args []string) {
	err := app.send(string(app.MenuItem.Type), fmt.Sprintf(app.MenuItem.Route, args[0], args[1]), map[string]string{})
	if err != nil {
		log.Printf("Recieve an error: %s", err.Error())
	}
}