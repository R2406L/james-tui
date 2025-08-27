package main

import (
	"fmt"
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

	err = g.SetKeybinding(header, 'e', gocui.ModNone, app.editConnection)
	if err != nil {
		return
	}

	err = g.SetKeybinding(body, gocui.KeyArrowUp, gocui.ModNone, moveBodyUp)
	if err != nil {
		return
	}

	err = g.SetKeybinding(body, gocui.KeyArrowDown, gocui.ModNone, moveBodyDown)
	if err != nil {
		return
	}

	err = g.SetKeybinding(menu, gocui.KeyArrowUp, gocui.ModNone, moveUp)
	if err != nil {
		return
	}

	err = g.SetKeybinding(menu, gocui.KeyArrowDown, gocui.ModNone, moveDown)
	if err != nil {
		return
	}

	err = g.SetKeybinding(menu, gocui.KeyEnter, gocui.ModNone, app.enter)
	if err != nil {
		return
	}

	err = g.SetKeybinding(editConnection, gocui.KeyEnter, gocui.ModNone, app.editConnectionSave)
	if err != nil {
		return
	}

	err = g.SetKeybinding(editConnection, gocui.KeyEsc, gocui.ModNone, app.editConnectionCancel)
	if err != nil {
		return
	}

	err = g.SetKeybinding(editConnection, gocui.KeyArrowLeft, gocui.ModNone, editorMoveLeft)
	if err != nil {
		return
	}

	err = g.SetKeybinding(editConnection, gocui.KeyArrowRight, gocui.ModNone, editorMoveRight)
	if err != nil {
		return
	}

	return
}

func (app *App) editConnection(g *gocui.Gui, v *gocui.View) error {
	view, err := g.View(editConnection)
	if err != nil || view == nil {
		return err
	}

	view.SetCursor(0, len(" Server address: "))
	view.Editable = true
	view.Visible = true
	app.View = view
	
	g.SetViewOnTop(editConnection)
	g.SetCurrentView(editConnection)

	return nil
}

func (app *App) editConnectionSave(g *gocui.Gui, v *gocui.View) error {
	app.View.Visible = false
	app.View.Editable = false
	app.Url = app.View.Buffer()

	headerView, err := g.View(header)
	if err != nil {
		return err
	}
	headerView.Title = fmt.Sprintf(titleHeader, app.Url)

	menuView, err := g.View(menu)
	if err != nil {
		return err
	}
	app.View = menuView

	g.SetCurrentView(menuView.Name())

	return nil
}

func (app *App) editConnectionCancel(g *gocui.Gui, v *gocui.View) error {
	app.View.Visible = false
	app.View.Editable = false
	
	menuView, err := g.View(menu)
	if err != nil {
		return err
	}
	
	app.View = menuView
	g.SetCurrentView(menuView.Name())
	
	return nil
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
		v.Clear()
	}
	

	if item.Route == "" {
		return nil
	}

	bodyView, err := g.View(body)
	if err != nil {
		app.Error = err.Error()
	}
	bodyView.Clear()

	err = app.send(string(item.Type), string(item.Route), map[string]string{})
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

func moveBodyUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		x, y := v.Origin()
		v.SetOrigin(x, y - 1)
	}
	return nil
}

func moveBodyDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		x, y := v.Origin()
		v.SetOrigin(x, y + 1)
	}
	return nil
}

func editorMoveUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		x, y := v.Cursor()
		v.SetCursor(x, y - 1)
	}
	return nil
}

func editorMoveDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		x, y := v.Cursor()
		v.SetCursor(x, y + 1)
	}
	return nil
}

func editorMoveLeft(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		x, y := v.Cursor()
		v.SetCursor(x - 1, y)
	}
	return nil
}

func editorMoveRight(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		x, y := v.Cursor()
		v.SetCursor(x + 1, y)
	}
	return nil
}

func (app *App) switchView(g *gocui.Gui, v *gocui.View) (err error) {
	newView, err := app.getNextView()
	if err != nil {
		return err
	}

	app.View = newView
	return
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}