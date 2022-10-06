package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func main() {
	fmt.Printf("running app!")
	application := app.New()
	window := application.NewWindow("Fyne App")
	window.Resize(fyne.NewSize(1000, 700))

	text1 := canvas.NewText("Hello", color.White)
	button := widget.NewButton("Click me!", func() {
		fmt.Printf("\n clicked")
		dialog.ShowInformation("Alert", "You clicked the button!", window)
	})

	textContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), text1, layout.NewSpacer())

	content := container.New(
		layout.NewHBoxLayout(),
		layout.NewSpacer(),
		container.New(
			layout.NewVBoxLayout(),
			layout.NewSpacer(),
			button,
			textContainer,
			layout.NewSpacer(),
		),
		layout.NewSpacer())

	window.SetContent(content)

	window.ShowAndRun()
}
