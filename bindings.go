package main

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
	"log"
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

	err = g.SetKeybinding(inputSimpleLayout_buttonOk, gocui.KeyEnter, gocui.ModNone, app.inputSimpleSave)
	if err != nil {
		return
	}

	err = g.SetKeybinding(inputSimpleLayout_buttonCancel, gocui.KeyEnter, gocui.ModNone, app.inputSimpleHide)
	if err != nil {
		return
	}

	err = g.SetKeybinding(inputEmailPasswordLayout_buttonOk, gocui.KeyEnter, gocui.ModNone, app.inputEmailPasswordLayoutSave)
	if err != nil {
		return
	}

	err = g.SetKeybinding(inputEmailPasswordLayout_buttonCancel, gocui.KeyEnter, gocui.ModNone, app.inputEmailPasswordHide)
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
		err := item.Function(g, v)
		if err != nil {
			log.Printf("Function error: %s", err.Error())
		}
		
		return nil
	}

	if item.Route == "" {
		return nil
	}

	err := app.send(string(item.Type), string(item.Route), map[string]string{})
	if err != nil {
		log.Printf("Server response error: %s", err.Error())
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
	view.Subtitle = subtitleInput
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

func (app *App) inputSimpleShow(g *gocui.Gui, v *gocui.View) (err error) {
	
	view, err := g.View(inputSimpleLayout)
	if err != nil || view == nil {
		return err
	}

	view_editor, err := g.View(inputSimpleLayout_editor)
	if err != nil || view_editor == nil {
		return err
	}

	view_ok, err := g.View(inputSimpleLayout_buttonOk)
	if err != nil || view == nil {
		return err
	}

	view_cancel, err := g.View(inputSimpleLayout_buttonCancel)
	if err != nil || view == nil {
		return err
	}

	view.Clear()
	view.Write([]byte(fmt.Sprintf("\n %s: ", app.MenuItem.Title)))
	view.Editable = false
	view.Visible = true
	g.SetViewOnTop(inputSimpleLayout)

	view_editor.Clear()
	view_editor.Editable = true
	view_editor.Visible = true
	app.View = view_editor
	g.SetViewOnTop(inputSimpleLayout_editor)
	g.SetCurrentView(inputSimpleLayout_editor)

	view_ok.Clear()
	view_ok.Editable = false
	view_ok.Visible = true
	view_ok.Highlight = false
	view_ok.Write([]byte("        OK"))
	g.SetViewOnTop(inputSimpleLayout_buttonOk)

	view_cancel.Clear()
	view_cancel.Editable = false
	view_cancel.Visible = true
	view_cancel.Highlight = false
	view_cancel.Write([]byte("      Cancel"))
	g.SetViewOnTop(inputSimpleLayout_buttonCancel)

	return nil
}

func (app *App) inputSimpleHide(g *gocui.Gui, v *gocui.View) (err error) {

	view, err := g.View(inputSimpleLayout)
	if err != nil || view == nil {
		return err
	}
	view.Visible = false

	view_editor, err := g.View(inputSimpleLayout_editor)
	if err != nil || view_editor == nil {
		return err
	}
	view_editor.Visible = false

	view_ok, err := g.View(inputSimpleLayout_buttonOk)
	if err != nil || view == nil {
		return err
	}
	view_ok.Visible = false

	view_cancel, err := g.View(inputSimpleLayout_buttonCancel)
	if err != nil || view == nil {
		return err
	}
	view_cancel.Visible = false

	app.selectMenuView(g, v)

	return nil
}

func (app *App) inputSimpleSave(g *gocui.Gui, v *gocui.View) error {
	
	view_editor, err := g.View(inputSimpleLayout_editor)
	if err != nil || view_editor == nil {
		return err
	}
	value := view_editor.Buffer()
	
	app.MenuItem.Handler([]string{value})

	app.inputSimpleHide(g, v)

	return nil
}

func (app *App) inputEmailPasswordShow(g *gocui.Gui, v *gocui.View) (err error) {

	view, err := g.View(inputEmailPasswordLayout)
	if err != nil || view == nil {
		return err
	}

	view_email, err := g.View(inputEmailPasswordLayout_email)
	if err != nil || view_email == nil {
		return err
	}

	view_password, err := g.View(inputEmailPasswordLayout_password)
	if err != nil || view_password == nil {
		return err
	}

	view_ok, err := g.View(inputEmailPasswordLayout_buttonOk)
	if err != nil || view == nil {
		return err
	}

	view_cancel, err := g.View(inputEmailPasswordLayout_buttonCancel)
	if err != nil || view == nil {
		return err
	}

	view.Title = app.MenuItem.Title
	view.Editable = false
	view.Visible = true
	g.SetViewOnTop(inputEmailPasswordLayout)

	view_email.Clear()
	view_email.Editable = true
	view_email.Visible = true
	app.View = view_email
	g.SetViewOnTop(inputEmailPasswordLayout_email)
	g.SetCurrentView(inputEmailPasswordLayout_email)

	view_password.Clear()
	view_password.Editable = true
	view_password.Visible = true
	view_password.Mask = '*'
	g.SetViewOnTop(inputEmailPasswordLayout_password)

	view_ok.Clear()
	view_ok.Editable = false
	view_ok.Visible = true
	view_ok.Highlight = false
	view_ok.Write([]byte("        OK"))
	g.SetViewOnTop(inputEmailPasswordLayout_buttonOk)

	view_cancel.Clear()
	view_cancel.Editable = false
	view_cancel.Visible = true
	view_cancel.Highlight = false
	view_cancel.Write([]byte("      Cancel"))
	g.SetViewOnTop(inputEmailPasswordLayout_buttonCancel)


	return nil
}

func (app *App) inputEmailPasswordHide(g *gocui.Gui, v *gocui.View) (err error) {

	view, err := g.View(inputEmailPasswordLayout)
	if err != nil || view == nil {
		return err
	}
	view.Visible = false

	view_email, err := g.View(inputEmailPasswordLayout_email)
	if err != nil || view_email == nil {
		return err
	}
	view_email.Visible = false

	view_password, err := g.View(inputEmailPasswordLayout_password)
	if err != nil || view_password == nil {
		return err
	}
	view_password.Visible = false

	view_ok, err := g.View(inputEmailPasswordLayout_buttonOk)
	if err != nil || view == nil {
		return err
	}
	view_ok.Visible = false

	view_cancel, err := g.View(inputEmailPasswordLayout_buttonCancel)
	if err != nil || view == nil {
		return err
	}
	view_cancel.Visible = false

	app.selectMenuView(g, v)

	return nil
}

func (app *App) inputEmailPasswordLayoutSave(g *gocui.Gui, v *gocui.View) (err error) {
	
	view_email, err := g.View(inputEmailPasswordLayout_email)
	if err != nil || view_email == nil {
		return err
	}
	email := view_email.Buffer()

	view_password, err := g.View(inputEmailPasswordLayout_password)
	if err != nil || view_password == nil {
		return err
	}
	password := view_password.Buffer()
	
	app.MenuItem.Handler([]string{email, password})

	app.inputEmailPasswordHide(g, v)

	return nil
}

func (app *App) selectMenuView(g *gocui.Gui, v *gocui.View) {

	menuView, err := g.View(menu)
	if err != nil {
		log.Printf("View %s not found", menu)
	}
	
	app.View = menuView
	g.SetCurrentView(menuView.Name())

}