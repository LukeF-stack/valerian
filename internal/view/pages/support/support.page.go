package support

import (
	"fmt"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"mobile-app/internal/types"
	"mobile-app/internal/view/context"
)

type Support struct {
	types.Page
}

func (support *Support) Init(c *context.Context) {
	support.Title = "Support"
	textBinding := binding.NewString()
	textBinding.Set(support.Title)
	text1 := widget.NewLabelWithData(textBinding)
	button := widget.NewButton("Click me!", func() {
		fmt.Printf("\n clicked")
		textBinding.Set("Changed")
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

	support.Content = content
}
