package context

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"mobile-app/internal/router"
)

type Context struct {
	App           fyne.App
	Route         *router.Router
	ActiveTasks   *fyne.Container
	CompleteTasks *fyne.Container
}

func (context *Context) Init(a *fyne.App, r *router.Router) {
	context.App = *a
	context.Route = r
	context.ActiveTasks = container.New(
		layout.NewVBoxLayout(),
	)
}

func (context *Context) NewNotify(title, content string) {
	notification := fyne.Notification{Title: title, Content: content}
	context.App.SendNotification(&notification)
}

func (context *Context) NewTask(title string) {
	//newTask := new(types.Task)
	//newTask.Complete = false
	//newTask.Name = title
	//newTask.Description = "This is a test task"

	context.ActiveTasks.Add(widget.NewLabel("test"))
}
