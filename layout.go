package main

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
	"io/ioutil"
	"log"
	"time"
)

func (app *App) mainLayout(g *gocui.Gui) (err error) {
	t := time.Now()
	log.Printf("-> Call under %v", t)

	maxX, maxY := g.Size()

	// Header layout
	v, err := g.SetView("header", int(0), int(0), maxX-1, maxY-1, byte(0))
	if err != nil && err != gocui.ErrUnknownView {
		return
	}
	v.Frame = true
	v.FrameRunes = []rune{'═','║','╔','╗','╚','╝'}
	v.Title = fmt.Sprintf(" %s [e]dit | Ctrl+C - [q]uit | Tab - Switch view ", app.Url)

	// Menu layout
	v, err = g.SetView("menu", int(1), int(1), int(30), maxY-2, byte(0))
	if err != nil && err != gocui.ErrUnknownView {
		return
	}

	v.Title = "Main menu"
	v.Highlight = true
	v.SelBgColor = gocui.ColorBlue
	v.SelFgColor = gocui.ColorWhite
	
	app.setMenu(v)
	
	// Body layout
	v, err = g.SetView("body", int(31), int(1), maxX-2, maxY-2, byte(0))
	if err != nil && err != gocui.ErrUnknownView {
		return
	}
	v.Title = "Server response"
	v.Wrap = true
	app.setResponse(g, v)
	
	g.SetCurrentView(app.View)

	return nil
}

func (app *App) setMenu(v *gocui.View) {
	if v != nil && len(v.Buffer()) == 0 {
		for _, e := range app.Menu.Elements {
			v.Write([]byte(fmt.Sprintf("%s\n", e.Name)))
		}
	}
}

func (app *App) setResponse(g *gocui.Gui, v *gocui.View) {
	v.Clear()
	view := g.CurrentView()
	if view != nil {
		x, y := view.Cursor()
		
		if len(app.Menu.Elements) > y {
			v.Write([]byte(fmt.Sprintf("Menu position: %s\n", app.Menu.Elements[y].Name)))
		}
		
		v.Write([]byte(fmt.Sprintf("Cursor position: %s (%d, %d)\n", view.Name(), x, y)))

		if app.Error != "" {
			v.Write([]byte(fmt.Sprintf("Server error: %s\n", app.Error)))
		}

		if app.Response != nil {
			v.Write([]byte(fmt.Sprintf("Server response: %d\n", app.Response.StatusCode)))
			v.Write([]byte(fmt.Sprintf("Server headers: %v\n", app.Response.Header)))

			body, _ := ioutil.ReadAll(app.Response.Body)
			v.Write([]byte("Server body:\n"))
			v.Write([]byte(fmt.Sprintf("%s\n", string(body))))
		}
	}
}