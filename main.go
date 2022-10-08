package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"mobile-app/internal/router"
	navigation "mobile-app/internal/view/components/nav"
	"mobile-app/internal/view/context"
	newPage "mobile-app/internal/view/pages/new"
	"mobile-app/internal/view/theme"
)

func main() {
	fmt.Printf("running app!")
	application := app.New()
	window := application.NewWindow("Fyne App")
	window.Resize(fyne.NewSize(500, 1000))
	application.Settings().SetTheme(new(theme.CustomTheme))
	route := new(router.Router)
	route.Init()

	ctx := new(context.Context)
	ctx.Init(&application, route)

	addPage := new(newPage.Add)
	addPage.Init(ctx)
	navBar := new(navigation.Nav)
	navBar.Init(ctx)

	title := widget.NewLabelWithData(route.Title)

	titleBox := container.New(layout.NewGridLayoutWithRows(1),
		container.New(
			layout.NewCenterLayout(),
			title,
		))

	nav := container.New(
		layout.NewPaddedLayout(),
		navBar.Render(),
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

	route.Render(addPage.Page)

	window.ShowAndRun()
}
