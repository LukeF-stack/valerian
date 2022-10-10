package context

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"mobile-app/internal/router"
	"mobile-app/internal/view/components/task"
)

type Context struct {
	App           fyne.App
	Route         *router.Router
	ActiveTasks   *TaskList
	CompleteTasks *TaskList
}

type TaskList struct {
	Container *fyne.Container
	List      []*task.Task
}

func (context *Context) Init(a *fyne.App, r *router.Router) {
	context.App = *a
	context.Route = r
	context.ActiveTasks = new(TaskList)
	context.ActiveTasks.Container = container.New(
		layout.NewVBoxLayout(),
	)
}

func (context *Context) NewNotify(title, content string) {
	notification := fyne.Notification{Title: title, Content: content}
	context.App.SendNotification(&notification)
}

func (context *Context) NewTask(title, description string) {
	newTask := task.CreateTask(title, description)
	context.ActiveTasks.List = append(context.ActiveTasks.List, newTask)
	context.ActiveTasks.Container.Add(newTask.Container)
}
