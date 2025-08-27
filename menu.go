package main

import (
	"github.com/awesome-gocui/gocui"
)

type Menu struct {
	Parent		*Menu
	Elements	[]*MenuItem
}

type MenuCallback func(g *gocui.Gui) (err error)

type MenuItem struct {
	Name		string
	Submenu		*Menu
	Route		string
	Type		string
	Title		string
	Function	MenuCallback
}

func (app *App) getMenu() *Menu {
	main := Menu{
		Elements: make([]*MenuItem, 0),
	}

	domain := Menu{
		Parent: &main,
		Elements: make([]*MenuItem, 0),
	}

	users := Menu{
		Parent: &main,
		Elements: make([]*MenuItem, 0),
	}

	domain.Elements = append(domain.Elements, &MenuItem{Name: "..", Submenu: &main, Route: "", Type: "",})
	domain.Elements = append(domain.Elements, &MenuItem{Name: "List", Submenu: nil, Route: "domains", Type: "GET",})
	domain.Elements = append(domain.Elements, &MenuItem{Name: "Check", Submenu: nil, Route: "domains/%s", Type: "GET", Title: "Enter domain name", Function: app.inputShow})
	domain.Elements = append(domain.Elements, &MenuItem{Name: "Add", Submenu: nil, Route: "domains/%s", Type: "PUT", Title: "Enter domain name", Function: app.inputShow})
	domain.Elements = append(domain.Elements, &MenuItem{Name: "Delete", Submenu: nil, Route: "domains/%s", Type: "DELETE", Title: "Enter domain name", Function: app.inputShow})

	users.Elements = append(users.Elements, &MenuItem{Name: "..", Submenu: &main, Route: "", Type: "",})
	users.Elements = append(users.Elements, &MenuItem{Name: "List", Submenu: nil, Route: "users", Type: "GET",})
	users.Elements = append(users.Elements, &MenuItem{Name: "Mailboxes", Submenu: nil, Route: "", Type: "",})
	users.Elements = append(users.Elements, &MenuItem{Name: "View", Submenu: nil, Route: "", Type: "",})
	users.Elements = append(users.Elements, &MenuItem{Name: "Add", Submenu: nil, Route: "", Type: "",})
	users.Elements = append(users.Elements, &MenuItem{Name: "Delete", Submenu: nil, Route: "", Type: "",})
	users.Elements = append(users.Elements, &MenuItem{Name: "Change password", Submenu: nil, Route: "", Type: "",})

	main.Elements = append(main.Elements, &MenuItem{Name: "HealthCheck", Submenu: nil, Route: "healthcheck", Type: "GET",})
	main.Elements = append(main.Elements, &MenuItem{Name: "Domains", Submenu: &domain, Route: "", Type: "",})
	main.Elements = append(main.Elements, &MenuItem{Name: "Users", Submenu: &users, Route: "", Type: "",})
	main.Elements = append(main.Elements, &MenuItem{Name: "Mailboxes", Submenu: nil, Route: "", Type: "",})
	main.Elements = append(main.Elements, &MenuItem{Name: "Messages", Submenu: nil, Route: "", Type: "",})
	main.Elements = append(main.Elements, &MenuItem{Name: "Quotas", Submenu: nil, Route: "", Type: "",})
	main.Elements = append(main.Elements, &MenuItem{Name: "Droplist", Submenu: nil, Route: "", Type: "",})
	main.Elements = append(main.Elements, &MenuItem{Name: "Queues", Submenu: nil, Route: "", Type: "",})
	main.Elements = append(main.Elements, &MenuItem{Name: "Sieve", Submenu: nil, Route: "", Type: "",})
	main.Elements = append(main.Elements, &MenuItem{Name: "Jmap", Submenu: nil, Route: "", Type: "",})
	main.Elements = append(main.Elements, &MenuItem{Name: "Tasks", Submenu: nil, Route: "", Type: "",})
	main.Elements = append(main.Elements, &MenuItem{Name: "Send email", Submenu: nil, Route: "", Type: "",})
	main.Elements = append(main.Elements, &MenuItem{Name: "Cassandra (*)", Submenu: nil, Route: "", Type: "",})

	main.Elements = append(main.Elements, &MenuItem{Name: "Settings", Submenu: nil, Route: "", Type: "",})
	
	return &main
}
