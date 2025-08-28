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

	err = g.SetKeybinding("", gocui.KeyF10, gocui.ModNone, quit)
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

	err = g.SetKeybinding(editConnection, gocui.KeyEsc, gocui.ModNone, app.cancel)
	if err != nil {
		return
	}

	err = g.SetKeybinding(editConnection, gocui.KeyArrowLeft, gocui.ModNone, moveLeft)
	if err != nil {
		return
	}

	err = g.SetKeybinding(editConnection, gocui.KeyArrowRight, gocui.ModNone, moveRight)
	if err != nil {
		return
	}

	err = g.SetKeybinding(inputLayout, gocui.KeyEnter, gocui.ModNone, app.inputSave)
	if err != nil {
		return
	}

	err = g.SetKeybinding(inputLayout, gocui.KeyEsc, gocui.ModNone, app.cancel)
	if err != nil {
		return
	}

	return
}

// Layout global function
func pass(g *gocui.Gui, v *gocui.View) error {
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func moveLeft(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		x, y := v.Cursor()
		v.SetCursor(x - 1, y)
	}
	return nil
}

func moveRight(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		x, y := v.Cursor()
		v.SetCursor(x + 1, y)
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

// Layout call function
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
	app.MenuItem = item
	
	if item.Function != nil {
		err := item.Function(g)
		if err != nil {
			app.Error = err.Error()
		}
		
		return nil
	}

	if item.Route == "" {
		return nil
	}

	err := app.send(string(item.Type), string(item.Route), map[string]string{})
	if err != nil {
		app.Error = err.Error()
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

func (app *App) cancel(g *gocui.Gui, v *gocui.View) error {
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

func (app *App) editConnection(g *gocui.Gui, v *gocui.View) error {
	view, err := g.View(editConnection)
	if err != nil || view == nil {
		return err
	}

	view.Title = titleEditConnection
	view.Write([]byte(app.Url))
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

func (app *App) inputShow(g *gocui.Gui) (err error) {
	view, err := g.View(inputLayout)
	if err != nil || view == nil {
		return err
	}
	view.Clear()

	view.Title = app.MenuItem.Title
	view.Subtitle = subtitleInput
	view.Editable = true
	view.Visible = true
	app.View = view
	
	g.SetViewOnTop(inputLayout)
	g.SetCurrentView(inputLayout)

	return nil
}

func (app *App) inputSave(g *gocui.Gui, v *gocui.View) error {
	app.View.Visible = false
	app.View.Editable = false
	
	value := app.View.Buffer()

	menuView, err := g.View(menu)
	if err != nil {
		return err
	}
	app.View = menuView

	g.SetCurrentView(menuView.Name())
	app.MenuItem.Handler(value)

	return nil
}