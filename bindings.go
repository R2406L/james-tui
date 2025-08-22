package main

import (
	"github.com/awesome-gocui/gocui"
)

func (app *App) keyBindings(g *gocui.Gui) (err error) {
	
	err = g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit)
	if err != nil {
		return
	}

	err = g.SetKeybinding("", 'q', gocui.ModNone, quit)
	if err != nil {
		return
	}

	err = g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, app.switchView)
	if err != nil {
		return
	}

	err = g.SetKeybinding("menu", gocui.KeyArrowUp, gocui.ModNone, moveUp)
	if err != nil {
		return
	}

	err = g.SetKeybinding("menu", gocui.KeyArrowDown, gocui.ModNone, moveDown)
	if err != nil {
		return
	}

	err = g.SetKeybinding("menu", gocui.KeyEnter, gocui.ModNone, app.enter)
	if err != nil {
		return
	}

	return
}

func (app *App) enter(g *gocui.Gui, v *gocui.View) error {
	_, y := v.Cursor()

	if y >= len(app.Menu.Elements) {
		v.Clear()
		app.Error = "Wrong menu element"
		return nil
	}

	item := app.Menu.Elements[y]	
	if app.Menu.Elements[y].Submenu != nil {
		app.Menu = item.Submenu
	}
	v.Clear()

	if item.Route == "" {
		return nil
	}

	err := app.send(string(item.Type), string(item.Route), map[string]string{})
	if err != nil {
		app.Error = err.Error()
	}

	return nil
}

func moveUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		x, y := v.Cursor()
		v.SetCursor(x, y - 1)
	}
	return nil
}

func moveDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		x, y := v.Cursor()
		v.SetCursor(x, y + 1)
	}
	return nil
}

func (app *App) switchView(g *gocui.Gui, v *gocui.View) (err error) {
	for i, j := range app.Views {
		if j != app.View {
			continue
		}
		
		if i < len(app.Views) - 1 {
			app.View = app.Views[i + 1]
		} else {
			app.View = app.Views[0]
		}

		break
	}

	return
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}