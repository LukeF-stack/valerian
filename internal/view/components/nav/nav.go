package navigation

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"mobile-app/internal/router"
	"mobile-app/internal/types"
	"mobile-app/internal/view/pages/about"
	"mobile-app/internal/view/pages/home"
	"mobile-app/internal/view/pages/support"
)

type Nav struct {
	types.Component
}

func (nav *Nav) Init(route *router.Router) {
	aboutPage := new(about.About)
	aboutPage.Init()
	homePage := new(home.Home)
	homePage.Init()
	supportPage := new(support.Support)
	supportPage.Init()
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
