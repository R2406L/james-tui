package main

import (
	"errors"
	"fmt"
	"github.com/awesome-gocui/gocui"
	"github.com/tidwall/pretty"
	"io/ioutil"
)

func (app *App) getNextView() (view *gocui.View, err error) {

	if len(app.Views) == 0 {
		return nil, errors.New("View list is empty")
	}

	found := 0
	for i, v := range app.Views {
		if v.Name() == app.View.Name() {
			found = i
			break
		}
	}

	if found == len(app.Views) {
		found = 0
	}

	for _, v := range app.Views[found + 1:len(app.Views)] {
		if v.Visible {
			return v, nil
		}
	}
	
	return app.Views[0], nil
}

func (app *App) defineWindows(g *gocui.Gui) (err error) {
	maxX, maxY := g.Size()

	// Header
	v, err := g.SetView("header", int(0), int(0), maxX-1, maxY-1, byte(0))
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Title = fmt.Sprintf(" %s [e]dit | Ctrl+C - [q]uit | Tab - Switch view ", app.Url)
	v.Frame = true
	v.FrameRunes = []rune{'═','║','╔','╗','╚','╝'}
	app.Views = append(app.Views, v)

	// Menu
	v, err = g.SetView("menu", int(1), int(1), int(30), maxY-2, byte(0))
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Title = "Main menu"
	v.Highlight = true
	v.SelBgColor = gocui.ColorBlue
	v.SelFgColor = gocui.ColorWhite
	app.Views = append(app.Views, v)
	app.View = v

	// Body
	v, err = g.SetView("body", int(31), int(1), maxX-2, maxY-2, byte(0))
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Title = "Server response"
	v.Wrap = true
	app.Views = append(app.Views, v)

	// Edit connection
	v, err = g.SetView("editConnection", maxX / 2 - 30, maxY / 2 - 1, maxX / 2 + 30, maxY / 2 + 1, byte(0))
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Title = "Edit connection"
	v.Highlight = true
	v.Frame = true
	v.Visible = false
	app.Views = append(app.Views, v)

	return nil
}

func (app *App) defineLayouts(g *gocui.Gui) (err error) {

	v, err := g.View("menu")
	if err != nil {
		return err
	}
	app.setMenu(v)

	v, err = g.View("body")
	if err != nil {
		return err
	}
	app.setResponse(g, v)

	v, err = g.View("editConnection")
	if err != nil {
		return err
	}
	app.setEditConnection(v)

	g.SetCurrentView(app.View.Name())

	return nil
}

func (app *App) setMenu(v *gocui.View) {
	if v != nil && len(v.Buffer()) == 0 {
		for _, e := range app.Menu.Elements {
			v.Write([]byte(fmt.Sprintf("%s\n", e.Name)))
		}
	}
}

func (app *App) setEditConnection(v *gocui.View) {
	if v != nil && len(v.Buffer()) == 0 {
		v.Write([]byte(fmt.Sprintf(" Server address: %s", app.Url)))
	}
}

func (app *App) setResponse(g *gocui.Gui, v *gocui.View) {
	if v != nil && len(v.Buffer()) == 0 && app.Response != nil {
		v.Write([]byte(fmt.Sprintf("Server response: %d\n", app.Response.StatusCode)))
		v.Write([]byte(fmt.Sprintf("Server headers: %v\n", app.Response.Header)))

		body, _ := ioutil.ReadAll(app.Response.Body)
		v.Write([]byte("Server body:\n"))
		v.Write([]byte(fmt.Sprintf("%s\n", pretty.Pretty(body))))
	}
}