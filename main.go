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
	View		*gocui.View
	Views		[]*gocui.View
	Response	*http.Response
	Error		string
}

func main() {
	app := App{
		Url: "http://127.0.0.1:8000/",
		Menu: getMenu(),
	}

	f, err := os.OpenFile("logfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
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

	for _, v := range app.Views {
		log.Print(v.Name())
	}

	err = g.MainLoop()
	if err != nil {
		log.Panic("Kernel error: %v", err)
	}
}
