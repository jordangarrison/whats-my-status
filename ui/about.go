package ui

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const version = "v0.3.2"

var releaseURL = &url.URL{
	Scheme: "https",
	Host:   "github.com",
	Path:   "/jordangarrison/whats-my-status/releases/tag/" + version,
}

type about struct {
	icon        *canvas.Image
	nameLabel   *widget.Label
	spacerLabel *widget.Label
	hyperlink   *widget.Hyperlink
}

func newAbout() *about {
	return &about{}
}

func (a *about) buildUI() *fyne.Container {
	a.icon = canvas.NewImageFromResource(theme.ComputerIcon())
	a.icon.SetMinSize(fyne.NewSize(256, 256))

	a.nameLabel = widget.NewLabel("Whats My Status?")
	a.spacerLabel = widget.NewLabel("-")
	a.hyperlink = &widget.Hyperlink{
		Text:      version,
		URL:       releaseURL,
		TextStyle: fyne.TextStyle{Bold: true},
	}

	return container.NewVBox(
		layout.NewSpacer(),
		container.NewHBox(layout.NewSpacer(), a.icon, layout.NewSpacer()),
		container.NewHBox(layout.NewSpacer(), a.nameLabel, a.spacerLabel, a.hyperlink, layout.NewSpacer()),
		layout.NewSpacer(),
	)
}

func (a *about) tabItem() *container.TabItem {
	return &container.TabItem{
		Text:    "About",
		Icon:    theme.InfoIcon(),
		Content: a.buildUI(),
	}
}
