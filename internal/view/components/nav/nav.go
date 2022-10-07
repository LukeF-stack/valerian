package navigation

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"mobile-app/internal/types"
	"mobile-app/internal/view/context"
	"mobile-app/internal/view/pages/about"
	"mobile-app/internal/view/pages/home"
	"mobile-app/internal/view/pages/support"
)

type Nav struct {
	types.Component
}

func (nav *Nav) Init(c *context.Context) {
	aboutPage := new(about.About)
	aboutPage.Init(c)
	homePage := new(home.Home)
	homePage.Init(c)
	supportPage := new(support.Support)
	supportPage.Init(c)
	route := c.Route
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
			widget.NewButton("About", func() { route.Render(aboutPage.Page) }),
			widget.NewButton("Home", func() { route.Render(homePage.Page) }),
			widget.NewButton("Support", func() { route.Render(supportPage.Page) }),
			layout.NewSpacer(),
		),
	)
}
