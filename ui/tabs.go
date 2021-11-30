package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	util "github.com/jordangarrison/whats-my-status/util"
)

// Create will stitch together all ui components
func Create(app fyne.App, window fyne.Window, config util.Config) *container.AppTabs {

	return &container.AppTabs{Items: []*container.TabItem{
		newStatusUI(app, window, config).tabItem(),
		newAliasUI(app, window, config).tabItem(),
		newAbout().tabItem(),
	}}
}
