package about

import (
	"fmt"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"mobile-app/internal/types"
)

type About struct {
	types.Page
}

func (about *About) Init() {
	about.Title = "About"
	textBinding := binding.NewString()
	textBinding.Set(about.Title)
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

	about.Content = content
}
