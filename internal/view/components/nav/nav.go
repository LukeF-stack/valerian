package navigation

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"mobile-app/internal/types"
	"mobile-app/internal/view/context"
	"mobile-app/internal/view/pages/active"
	"mobile-app/internal/view/pages/complete"
	newPage "mobile-app/internal/view/pages/new"
)

type Nav struct {
	types.Component
}

func (nav *Nav) Init(c *context.Context) {
	activePage := new(active.Active)
	activePage.Init(c)
	addPage := new(newPage.Add)
	addPage.Init(c)
	completePage := new(complete.Complete)
	completePage.Init(c)
	route := c.Route
	resource, err := fyne.LoadResourceFromPath("internal/resources/add.svg")
	if err != nil {
		fmt.Printf(err.Error())
	}
	addBtn := widget.NewButtonWithIcon("New", resource, func() { route.Render(addPage.Page) })
	nav.Content = container.New(
		layout.NewBorderLayout(
			layout.NewSpacer(),
			layout.NewSpacer(),
			nil,
			nil,
		),
		layout.NewSpacer(),
		layout.NewSpacer(),
		container.New(
			layout.NewGridLayout(5),
			layout.NewSpacer(),
			widget.NewButton("Active", func() { route.Render(activePage.Page) }),
			addBtn,
			widget.NewButton("Completed", func() { route.Render(completePage.Page) }),
			layout.NewSpacer(),
		),
	)
}
