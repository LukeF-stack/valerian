package context

import (
	"fyne.io/fyne/v2"
	"mobile-app/internal/router"
)

type Context struct {
	App   fyne.App
	Route *router.Router
}

func (context *Context) Init(a *fyne.App, r *router.Router) {
	context.App = *a
	context.Route = r
}

func (context *Context) NewNotify(title, content string) {
	notification := fyne.Notification{Title: title, Content: content}
	context.App.SendNotification(&notification)
}
