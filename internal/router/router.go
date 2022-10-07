package router

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"mobile-app/internal/view/pages/home"
)

type Router struct {
	Content *fyne.Container
	Title   binding.String
}

func (router *Router) Init() {
	router.Title = binding.NewString()
	router.Title.Set("")
	router.Content = container.New(
		layout.NewMaxLayout(),
	)
}

func (router *Router) Render(component home.Home) {
	fmt.Printf(component.Title)
	router.Content.RemoveAll()
	router.Title.Set(component.Title)
	router.Content.Add(component.Content)
}
