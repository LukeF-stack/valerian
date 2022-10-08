package types

import "fyne.io/fyne/v2"

type Componenter interface {
	Init()
	Render()
}

type Component struct {
	Componenter
	Content *fyne.Container
}

type Page struct {
	Component
	Title string
}

func (component *Component) Render() *fyne.Container {
	return component.Content
}

type Task struct {
	Name        string
	Description string
	Complete    bool
}
