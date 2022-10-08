package active

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"mobile-app/internal/types"
	"mobile-app/internal/view/context"
)

type Active struct {
	types.Page
	activeTasks *fyne.Container
}

func (active *Active) Init(c *context.Context) {
	active.Title = "Active Tasks"

	content := container.NewMax(
		container.NewVScroll(
			c.ActiveTasks,
		),
	)

	active.Content = content
}
