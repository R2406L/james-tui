package main

import (
	"github.com/awesome-gocui/gocui"
	"log"
	"net/http"
	"os"
)

type App struct {
	Url			string
	Menu		*Menu
	MenuItem	*MenuItem
	View		*gocui.View
	Views		[]*gocui.View
	Response	*http.Response
	Error		string
}

const (
	body = "body"
	header = "header"
	menu = "menu"
	editConnection = "editConnection"
	inputLayout = "inputLayout"
	titleHeader = " %s [e]dit | F10 - quit | Tab - Switch view "
	titleMenu = "Main menu"
	titleBody = "Server response"
	titleEditConnection = "Edit connection to server"
	titleEnterEmail = "Edit email address"
	subtitleInput = "Enter - Save | Esc - Cancel"
)

func main() {

	app := App{
		Url: "http://127.0.0.1:8000/",
	}
	app.Menu = app.getMenu()

	f, err := os.OpenFile("james_tui.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
    	log.Fatalf("Error opening log file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	g, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Cursor = true
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen
	g.SelFrameColor = gocui.ColorGreen
	
	g.SetManagerFunc(app.defineLayouts)
	err = app.defineWindows(g)
	if err != nil {
		log.Fatalf("Error set layout settings: %v", err)
	}
		
	err = app.keyBindings(g)
	if err != nil {
		log.Fatalf("Error key bindings: %v", err)
	}

	err = g.MainLoop()
	if err != nil {
		log.Panic("Kernel error: %v", err)
	}
}
