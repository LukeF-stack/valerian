package complete

import (
	"fmt"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"mobile-app/internal/types"
	"mobile-app/internal/view/context"
)

type Complete struct {
	types.Page
}

func (complete *Complete) Init(c *context.Context) {
	complete.Title = "Complete Tasks"
	textBinding := binding.NewString()
	textBinding.Set(complete.Title)
	text1 := widget.NewLabelWithData(textBinding)
	button := widget.NewButton("Click me!", func() {
		fmt.Printf("\n clicked")
		for _, task := range c.ActiveTasks.List {
			task.Name.Set("changed")
		}
	})

	textContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), text1, layout.NewSpacer())

	content := container.New(
		layout.NewHBoxLayout(),
		layout.NewSpacer(),
		container.New(
			layout.NewVBoxLayout(),
			layout.NewSpacer(),
			textContainer,
			button,
			layout.NewSpacer(),
		),
		layout.NewSpacer())

	complete.Content = content
}
