package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/jordangarrison/whats-my-status/status"
	util "github.com/jordangarrison/whats-my-status/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	config  util.Config
)

func main() {
	// Config
	err := viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	// Set up the Application
	wms := app.New()
	window := wms.NewWindow("What's My Status?")
	window.CenterOnScreen()
	window.Resize(fyne.NewSize(800, 600))

	// Top Container
	statusText := widget.NewEntry()
	statusText.SetPlaceHolder("Enter your status message here...")
	topContainer := container.New(layout.NewCenterLayout(), statusText)

	// Set status and clear status buttons
	setStatusButton := widget.NewButton("Set Status", setStatus)
	clearStatusButton := widget.NewButton("Clear Status", clearStatus)

	// Bottom Container
	bottomContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), setStatusButton, clearStatusButton, layout.NewSpacer())
	window.SetContent(container.New(layout.NewVBoxLayout(), topContainer, layout.NewSpacer(), bottomContainer))

	// Run the Application
	window.ShowAndRun()
}

func setStatus() {
	fmt.Println("Setting status")
}

func clearStatus() {
	fmt.Println("Clearing status")
	// Find status alias
	for _, alias := range config.Aliases {
		if alias.Name == "clear" {
			config.Status = alias.Status
			break
		}
	}

	// Set status
	err := status.SetStatus(config)
	if err != nil {
		fmt.Println(err)
	}
}

func init() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".whats-my-status" (without extension).
		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.AddConfigPath(fmt.Sprintf("%s/.config", home))
		viper.SetConfigType("yaml")
		viper.SetConfigName(".wms.yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
