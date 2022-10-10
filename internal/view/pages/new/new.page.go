package new

import (
	"fmt"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"mobile-app/internal/types"
	"mobile-app/internal/view/context"
)

type Add struct {
	types.Page
}

func (add *Add) Init(c *context.Context) {
	add.Title = "Home"
	textBinding := binding.NewString()
	textBinding.Set(add.Title)
	text1 := widget.NewLabelWithData(textBinding)
	button := widget.NewButton("Click me!", func() {
		fmt.Printf("\n clicked")
		c.NewTask("test", "this is a test")
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

	add.Content = content
}
