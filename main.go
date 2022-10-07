package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"mobile-app/internal/view/pages/home"
)

func main() {
	fmt.Printf("running app!")
	application := app.New()
	window := application.NewWindow("Fyne App")
	window.Resize(fyne.NewSize(1000, 700))

	h := new(home.Home)
	h.Init()

	subContent := container.New(layout.NewHBoxLayout(), widget.NewButton("test", nil), widget.NewButton("test", nil), widget.NewButton("test", nil))

	content := container.New(
		layout.NewVBoxLayout(),
		subContent,
	)

	//tabs := container.NewAppTabs(
	//	container.NewTabItem("Tab 1", widget.NewLabel("Hello")),
	//	container.NewTabItemWithIcon("Home", theme.HomeIcon(), h.Content),
	//	container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	//)
	//
	//tabs.SetTabLocation(container.TabLocationBottom)
	//
	//tabs.Resize(fyne.NewSize(100, 100))

	window.SetContent(content)

	window.ShowAndRun()
}
