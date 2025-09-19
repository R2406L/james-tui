package main

import (
	"github.com/awesome-gocui/gocui"
)

type Menu struct {
	Parent		*Menu
	Elements	[]*MenuItem
}

type MenuCallback func(g *gocui.Gui, v *gocui.View) (err error)
type MenuHandler func(args []string)

type MenuItem struct {
	Name		string
	Submenu		*Menu
	Route		string
	Type		string
	Title		string
	Body		string
	Function	MenuCallback
	Handler		MenuHandler
}

func (app *App) getMenu() *Menu {
	main := Menu{
		Elements: make([]*MenuItem, 0),
	}

	domain := Menu{
		Parent: &main,
		Elements: make([]*MenuItem, 0),
	}

	droplist := Menu{
		Parent: &main,
		Elements: make([]*MenuItem, 0),
	}

	users := Menu{
		Parent: &main,
		Elements: make([]*MenuItem, 0),
	}

	mailboxes := Menu{
		Parent: &main,
		Elements: make([]*MenuItem, 0),
	}

	messages := Menu{
		Parent: &main,
		Elements: make([]*MenuItem, 0),
	}

	tasks := Menu{
		Parent: &main,
		Elements: make([]*MenuItem, 0),
	}

	tasks_status := Menu{
		Parent: &main,
		Elements: make([]*MenuItem, 0),
	}

	sieve := Menu{
		Parent: &main,
		Elements: make([]*MenuItem, 0),
	}

	jmap := Menu{
		Parent: &main,
		Elements: make([]*MenuItem, 0),
	}

	quotas := Menu{
		Parent: &main,
		Elements: make([]*MenuItem, 0),
	}

	queues := Menu{
		Parent: &main,
		Elements: make([]*MenuItem, 0),
	}

	domain.Elements = append(domain.Elements, &MenuItem{Name: "..", Submenu: &main, Route: "", Type: "",})
	domain.Elements = append(domain.Elements, &MenuItem{Name: "List", Submenu: nil, Route: "domains", Type: "GET",})
	domain.Elements = append(domain.Elements, &MenuItem{Name: "Check", Submenu: nil, Route: "domains/%s", Type: "GET", Title: "Enter domain name", Function: app.inputSimpleShow, Handler: app.simpleHandler})
	domain.Elements = append(domain.Elements, &MenuItem{Name: "Add", Submenu: nil, Route: "domains/%s", Type: "PUT", Title: "Enter domain name", Function: app.inputSimpleShow, Handler: app.simpleHandler})
	domain.Elements = append(domain.Elements, &MenuItem{Name: "Delete", Submenu: nil, Route: "domains/%s", Type: "DELETE", Title: "Enter domain name", Function: app.inputSimpleShow, Handler: app.simpleHandler})

	users.Elements = append(users.Elements, &MenuItem{Name: "..", Submenu: &main, Route: "", Type: "",})
	users.Elements = append(users.Elements, &MenuItem{Name: "List", Submenu: nil, Route: "users", Type: "GET",})
	users.Elements = append(users.Elements, &MenuItem{Name: "Mailboxes", Submenu: nil, Route: "users/%s/mailboxes", Type: "GET", Title: "Enter email address", Function: app.inputSimpleShow, Handler: app.simpleHandler})
	users.Elements = append(users.Elements, &MenuItem{Name: "View", Submenu: nil, Route: "users/%s", Type: "GET", Title: "Enter email address", Function: app.inputSimpleShow, Handler: app.simpleHandler})
	users.Elements = append(users.Elements, &MenuItem{Name: "Add", Submenu: nil, Route: "users/%s", Type: "PUT", Title: "Enter email address", Function: app.inputEmailPasswordLayoutShow, Handler: app.addUserHandler})
	users.Elements = append(users.Elements, &MenuItem{Name: "Delete", Submenu: nil, Route: "users/%s", Type: "DELETE", Title: "Enter email address", Function: app.inputSimpleShow, Handler: app.simpleHandler})
	users.Elements = append(users.Elements, &MenuItem{Name: "Change password", Submenu: nil, Route: "users/%s?force", Type: "PUT", Title: "Enter password", Function: app.inputEmailPasswordLayoutShow, Handler: app.changePasswordHandler})

	mailboxes.Elements = append(mailboxes.Elements, &MenuItem{Name: "..", Submenu: &main, Route: "", Type: "",})
	mailboxes.Elements = append(mailboxes.Elements, &MenuItem{Name: "Reindex all", Submenu: nil, Route: "mailboxes?task=reIndex", Type: "POST",})
	mailboxes.Elements = append(mailboxes.Elements, &MenuItem{Name: "Clear mailbox", Submenu: nil, Route: "users/%s/mailboxes/%s/messages", Type: "DELETE", Title: "Clear user mailbox", Body: "\n Enter email: \n\n\n Enter mailbox: \n", Function: app.inputDuoLayoutShow, Handler: app.mailboxClearHandler})

	tasks.Elements = append(tasks.Elements, &MenuItem{Name: "..", Submenu: &main, Route: "", Type: "",})
	tasks.Elements = append(tasks.Elements, &MenuItem{Name: "List", Submenu: &tasks_status, Route: "", Type: "",})
	tasks.Elements = append(tasks.Elements, &MenuItem{Name: "Detail", Submenu: nil, Route: "tasks/%s", Type: "GET", Title: "Enter task ID", Function: app.inputSimpleShow, Handler: app.simpleHandler})

	tasks_status.Elements = append(tasks_status.Elements, &MenuItem{Name: "..", Submenu: &tasks, Route: "", Type: "",})
	tasks_status.Elements = append(tasks_status.Elements, &MenuItem{Name: "All", Submenu: nil, Route: "tasks", Type: "GET",})
	tasks_status.Elements = append(tasks_status.Elements, &MenuItem{Name: "Waiting", Submenu: nil, Route: "tasks?status=waiting", Type: "GET",})
	tasks_status.Elements = append(tasks_status.Elements, &MenuItem{Name: "InProgress", Submenu: nil, Route: "tasks?status=inProgress", Type: "GET",})
	tasks_status.Elements = append(tasks_status.Elements, &MenuItem{Name: "Cancelled", Submenu: nil, Route: "tasks?status=canceled", Type: "GET",})
	tasks_status.Elements = append(tasks_status.Elements, &MenuItem{Name: "Completed", Submenu: nil, Route: "tasks?status=completed", Type: "GET",})
	tasks_status.Elements = append(tasks_status.Elements, &MenuItem{Name: "Failed", Submenu: nil, Route: "tasks?status=failed", Type: "GET",})

	queues.Elements = append(queues.Elements, &MenuItem{Name: "..", Submenu: &main, Route: "", Type: "",})
	queues.Elements = append(queues.Elements, &MenuItem{Name: "All", Submenu: nil, Route: "mailQueues", Type: "GET",})
	queues.Elements = append(queues.Elements, &MenuItem{Name: "Detail", Submenu: nil, Route: "mailQueues/%s", Type: "GET", Title: "Enter queue", Function: app.inputSimpleShow, Handler: app.simpleHandler,})
	queues.Elements = append(queues.Elements, &MenuItem{Name: "Messages", Submenu: nil, Route: "mailQueues/%s/mails", Type: "GET", Title: "Enter queue", Function: app.inputSimpleShow, Handler: app.simpleHandler,})
	queues.Elements = append(queues.Elements, &MenuItem{Name: "Flush", Submenu: nil, Route: "mailQueues/%s?delayed=true", Type: "PATCH", Title: "Enter queue", Function: app.inputSimpleShow, Handler: app.simpleHandler,})
	queues.Elements = append(queues.Elements, &MenuItem{Name: "Update browse start", Submenu: nil, Route: "/mailQueues/%s?action=updateBrowseStart", Type: "POST", Title: "Enter queue", Function: app.inputSimpleShow, Handler: app.simpleHandler,})

	messages.Elements = append(messages.Elements, &MenuItem{Name: "..", Submenu: &main, Route: "", Type: "",})

	sieve.Elements = append(droplist.Elements, &MenuItem{Name: "..", Submenu: &main, Route: "", Type: "",})

	droplist.Elements = append(droplist.Elements, &MenuItem{Name: "..", Submenu: &main, Route: "", Type: "",})

	jmap.Elements = append(jmap.Elements, &MenuItem{Name: "..", Submenu: &main, Route: "", Type: "",})

	quotas.Elements = append(quotas.Elements, &MenuItem{Name: "..", Submenu: &main, Route: "", Type: "",})

	main.Elements = append(main.Elements, &MenuItem{Name: "HealthCheck", Submenu: nil, Route: "healthcheck", Type: "GET",})
	main.Elements = append(main.Elements, &MenuItem{Name: "Domains", Submenu: &domain, Route: "", Type: "",})
	main.Elements = append(main.Elements, &MenuItem{Name: "Users", Submenu: &users, Route: "", Type: "",})
	main.Elements = append(main.Elements, &MenuItem{Name: "Mailboxes", Submenu: &mailboxes, Route: "", Type: "",})
	main.Elements = append(main.Elements, &MenuItem{Name: "Messages", Submenu: &messages, Route: "", Type: "",})
	main.Elements = append(main.Elements, &MenuItem{Name: "Quotas", Submenu: &quotas, Route: "", Type: "",})
	main.Elements = append(main.Elements, &MenuItem{Name: "Droplist", Submenu: &droplist, Route: "", Type: "",})
	main.Elements = append(main.Elements, &MenuItem{Name: "Queues", Submenu: &queues, Route: "", Type: "",})
	main.Elements = append(main.Elements, &MenuItem{Name: "Sieve", Submenu: &sieve, Route: "", Type: "",})
	main.Elements = append(main.Elements, &MenuItem{Name: "Jmap", Submenu: &jmap, Route: "", Type: "",})
	main.Elements = append(main.Elements, &MenuItem{Name: "Tasks", Submenu: &tasks, Route: "", Type: "",})
	main.Elements = append(main.Elements, &MenuItem{Name: "Send email", Submenu: nil, Route: "", Type: "",})
	
	return &main
}
