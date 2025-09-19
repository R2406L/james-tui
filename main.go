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

	version = "v0.0.3"

	body = "body"
	header = "header"
	menu = "menu"
	notice = "notice"
	editConnection = "editConnection"
	inputLayout = "inputLayout"

	inputSimpleLayout = "inputSimpleLayout"
	inputSimpleLayout_editor = "inputSimpleLayout_email"
	inputSimpleLayout_buttonCancel = "inputSimpleLayout_buttonCancel"
	inputSimpleLayout_buttonOk = "inputSimpleLayout_buttonOk"

	inputEmailPasswordLayout = "inputEmailPasswordLayout"
	inputEmailPasswordLayout_email = "inputEmailPasswordLayout_email"
	inputEmailPasswordLayout_password = "inputEmailPasswordLayout_password"
	inputEmailPasswordLayout_buttonCancel = "inputEmailPasswordLayout_buttonCancel"
	inputEmailPasswordLayout_buttonOk = "inputEmailPasswordLayout_buttonOk"
	inputEmailPasswordLayout_body = "\n Enter email: \n\n\n Enter password: \n"
	
	inputDuoLayout = "inputDuoLayout"
	inputDuoLayout_first = "inputDuoLayout_first"
	inputDuoLayout_second = "inputDuoLayout_second"
	inputDuoLayout_buttonCancel = "inputDuoLayout_buttonCancel"
	inputDuoLayout_buttonOk = "inputDuoLayout_buttonOk"
	inputDuoLayout_body = "\n Enter first arg: \n\n\n Enter second arg: \n"

	titleHeader = " %s "
	titleHeaderSubtitle = " [e]dit | F10 - quit | Tab - Switch view ═ Version: %s "
	titleMenu = " Main menu "
	titleBody = " Server response "
	titleEditConnection = " Edit connection to server "
	titleEnterEmail = " Edit email address "
	titleNotice = " Notice "
	subtitleInput = " Enter - Save | Esc - Cancel "
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
