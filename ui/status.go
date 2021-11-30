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

type StatusUI struct {
	setStatusButton   *widget.Button
	clearStatusButton *widget.Button
	emoji             *widget.Entry
	statusText        *widget.Entry
	statusTime        *widget.Entry
	window            fyne.Window
	app               fyne.App
	config            util.Config
}

func newStatusUI(app fyne.App, window fyne.Window, config util.Config) *StatusUI {
	return &StatusUI{app: app, window: window, config: config}
}

func (sui *StatusUI) buildUI() *fyne.Container {
	sui.setStatusButton = widget.NewButton("Set Status", sui.setStatus)
	sui.clearStatusButton = widget.NewButton("Clear Status", sui.clearStatus)
	sui.statusText = widget.NewEntry()
	sui.statusText.SetPlaceHolder("Enter your status message here...")
	sui.emoji = widget.NewEntry()
	sui.emoji.SetPlaceHolder("Enter your emoji here... (e.g. :speech_balloon:)")
	sui.statusTime = widget.NewEntry()
	sui.statusTime.SetPlaceHolder("Enter your status time here... (e.g. 30m)")

	buttonContainer := container.NewHBox(sui.setStatusButton, sui.clearStatusButton)
	box := container.NewVBox(sui.statusText, sui.emoji, sui.statusTime, buttonContainer)
	return box
}

func (sui *StatusUI) tabItem() *container.TabItem {
	return &container.TabItem{
		Text:    "Set Status",
		Icon:    theme.ComputerIcon(),
		Content: sui.buildUI(),
	}
}

func (sui *StatusUI) setStatus() {
	fmt.Println("Setting status " + sui.statusText.Text + " " + sui.emoji.Text + " for " + sui.statusTime.Text)
	sui.config.Status = util.Status{
		StatusMessage: sui.statusText.Text,
		Emoji:         sui.emoji.Text,
		Time:          sui.statusTime.Text,
	}

	// Get the time
	if sui.config.Status.Time != "" {
		// epoch
		epoch, err := util.GetEpochTime(sui.config.Status.Time)
		if err != nil {
			fmt.Printf("Error getting epoch: %v", err)
		}
		sui.config.Status.Epoch = epoch

		iso8601, err := util.GetISO8601Time(sui.config.Status.Time)
		if err != nil {
			fmt.Printf("Error getting ISO8601: %v", err)
		}
		sui.config.Status.ISO8601 = iso8601
	}

	// Set the status
	err := status.SetStatus(sui.config)
	if err != nil {
		fmt.Printf("Error setting status: %v", err)
	}
}

func (sui *StatusUI) clearStatus() {
	fmt.Println("Clearing status")
	// Find status alias
	alias := sui.getAlias("clear")
	sui.config.Status = alias.Status

	// Set status
	err := status.SetStatus(sui.config)
	if err != nil {
		fmt.Println(err)
	}

	// Clear the status
	sui.statusText.SetText("")
	sui.emoji.SetText("")
	sui.statusTime.SetText("")
}

func (sui *StatusUI) getAlias(name string) *util.Alias {
	fmt.Println("Getting alias: " + name)
	for _, alias := range sui.config.Aliases {
		if alias.Name == name {
			fmt.Printf("Got alias: %+v\n", alias)
			return &alias
		}
	}
	return nil
}
