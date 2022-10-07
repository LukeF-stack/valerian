package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"mobile-app/internal/router"
	"mobile-app/internal/view/pages/home"
)

func main() {
	fmt.Printf("running app!")
	application := app.New()
	window := application.NewWindow("Fyne App")
	window.Resize(fyne.NewSize(500, 1000))
	homepage := new(home.Home)
	homepage.Init()
	route := new(router.Router)
	route.Init()

	title := widget.NewLabelWithData(r.Title)

	titleBox := container.New(layout.NewGridLayoutWithRows(1),
		container.New(
			layout.NewCenterLayout(),
			title,
		))

	nav := container.New(
		layout.NewPaddedLayout(),
		container.New(
			layout.NewGridLayout(3),
			widget.NewButton("About", nil),
			widget.NewButton("Home", nil),
			widget.NewButton("Support", nil),
		),
	)

	window.SetContent(
		container.New(
			layout.NewBorderLayout(
				titleBox,
				nav,
				nil,
				nil,
			),
			titleBox,
			nav,
			container.New(
				layout.NewAdaptiveGridLayout(1),
				route.Content,
			),
		),
	)

	route.Render(*homepage)

	window.ShowAndRun()
}
