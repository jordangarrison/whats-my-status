package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	ui "github.com/jordangarrison/whats-my-status/ui"
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
	config.Aliases = config.GetStatusAliases()

	// Set up the Application
	wms := app.NewWithID("io.github.jordangarrison.whats-my-status")
	window := wms.NewWindow("What's My Status?")
	window.CenterOnScreen()
	window.Resize(fyne.NewSize(800, 600))

	// Set up the window
	window.SetContent(ui.Create(wms, window, config))
	window.SetMaster()

	// Run the Application
	window.ShowAndRun()
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
