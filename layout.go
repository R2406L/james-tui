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
	v, err := g.SetView(header, int(0), int(0), maxX-1, maxY-1, byte(0))
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Title = fmt.Sprintf(titleHeader, app.Url)
	v.Frame = true
	v.FrameRunes = []rune{'═','║','╔','╗','╚','╝'}
	app.Views = append(app.Views, v)

	// Menu
	v, err = g.SetView(menu, int(1), int(1), int(30), maxY-10, byte(0))
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Title = titleMenu
	v.Highlight = true
	v.SelBgColor = gocui.ColorBlue
	v.SelFgColor = gocui.ColorWhite
	app.Views = append(app.Views, v)
	app.View = v

	// Notice window
	v, err = g.SetView(notice, int(1), maxY-9, int(30), maxY-2, byte(0))
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Title = titleNotice
	v.Highlight = false
	app.Views = append(app.Views, v)

	// Body
	v, err = g.SetView(body, int(31), int(1), maxX-2, maxY-2, byte(0))
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Title = titleBody
	v.Highlight = false
	v.Wrap = true
	app.Views = append(app.Views, v)

	// Edit connection
	v, err = g.SetView(editConnection, maxX / 2 - 30, maxY / 2 - 1, maxX / 2 + 30, maxY / 2 + 1, byte(0))
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Highlight = false
	v.Frame = true
	v.Visible = false
	app.Views = append(app.Views, v)

	// Input Layout
	v, err = g.SetView(inputLayout, maxX / 2 - 30, maxY / 2 - 1, maxX / 2 + 30, maxY / 2 + 1, byte(0))
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Highlight = false
	v.Frame = true
	v.Visible = false
	app.Views = append(app.Views, v)

	// Input multiple layots
	v, err = g.SetView(inputEmailPasswordLayout, maxX / 2 - 30, maxY / 2 - 10, maxX / 2 + 30, maxY / 2 + 1, byte(0))
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Highlight = false
	v.Frame = true
	v.Visible = false
	v.Write([]byte(inputEmailPasswordLayout_body))
	app.Views = append(app.Views, v)

	v, err = g.SetView(inputEmailPasswordLayout_email, maxX / 2 - 8, maxY / 2 - 9, maxX / 2 + 28, maxY / 2 - 7, byte(0))
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Highlight = false
	v.Frame = true
	v.Visible = false
	app.Views = append(app.Views, v)

	v, err = g.SetView(inputEmailPasswordLayout_password, maxX / 2 - 8, maxY / 2 - 6, maxX / 2 + 28, maxY / 2 - 4, byte(0))
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Highlight = false
	v.Frame = true
	v.Visible = false
	app.Views = append(app.Views, v)

	v, err = g.SetView(inputEmailPasswordLayout_buttonCancel, maxX / 2 - 25, maxY / 2 - 2, maxX / 2 - 5, maxY / 2, byte(0))
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Highlight = false
	v.Frame = true
	v.Visible = false
	app.Views = append(app.Views, v)

	v, err = g.SetView(inputEmailPasswordLayout_buttonOk, maxX / 2 + 5, maxY / 2 - 2, maxX / 2 + 25, maxY / 2, byte(0))
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Highlight = false
	v.Frame = true
	v.Visible = false
	app.Views = append(app.Views, v)


	return nil

}

func (app *App) defineLayouts(g *gocui.Gui) (err error) {

	v, err := g.View(menu)
	if err != nil {
		return err
	}
	
	if len(v.Buffer()) == 0 {
		for _, e := range app.Menu.Elements {
			v.Write([]byte(fmt.Sprintf("%s\n", e.Name)))
		}
	}

	v, err = g.View(notice)
	if err != nil {
		return err
	}
	app.setNotice(v)

	g.SetCurrentView(app.View.Name())

	return nil

}

func (app *App) setResponse() {
	
	var view *gocui.View
	for _, v := range app.Views {
		if v.Name() == body {
			view = v
			break
		}
	}
	view.Clear()

	if view != nil && app.Response != nil {
		body, _ := ioutil.ReadAll(app.Response.Body)
		view.Write([]byte(fmt.Sprintf("%s\n", pretty.Pretty(body))))
	}

}

func (app *App) setNotice(v *gocui.View) {
	
	v.Clear()

	if v != nil && app.Response != nil {
		v.Write([]byte(fmt.Sprintf("Server response code: %d\n", app.Response.StatusCode)))
	}

}