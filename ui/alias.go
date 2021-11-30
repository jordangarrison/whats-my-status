package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/jordangarrison/whats-my-status/status"
	util "github.com/jordangarrison/whats-my-status/util"
)

type AliasUI struct {
	app          fyne.App
	window       fyne.Window
	config       util.Config
	aliasButtons []*widget.Button
}

func newAliasUI(app fyne.App, window fyne.Window, config util.Config) *AliasUI {
	return &AliasUI{app: app, window: window, config: config}
}

func (aui *AliasUI) tabItem() *container.TabItem {
	return &container.TabItem{
		Text:    "Aliases",
		Icon:    theme.ListIcon(),
		Content: aui.buildUI(),
	}
}

func (aui *AliasUI) buildUI() *fyne.Container {
	aui.aliasButtons = make([]*widget.Button, len(aui.config.Aliases))
	grid := container.NewGridWrap(fyne.NewSize(110, 60))
	for i, alias := range aui.config.Aliases {
		button := widget.NewButton(alias.Name, aui.setAlias(alias.Name))
		button.SetText(alias.Name)
		button.Resize(fyne.NewSize(100, 50))
		aui.aliasButtons[i] = button
		grid.Add(aui.aliasButtons[i])
	}
	return grid
}

func (aui *AliasUI) setAlias(name string) func() {
	return func() {
		fmt.Println("Setting alias: " + name)
		alias := aui.getAlias(name)
		aui.config.Status = alias.Status
		err := status.SetStatus(aui.config)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (aui *AliasUI) getAlias(name string) *util.Alias {
	for _, alias := range aui.config.Aliases {
		if alias.Name == name {
			fmt.Printf("Found alias: %v\n", alias)
			return &alias
		}
	}
	return nil
}
