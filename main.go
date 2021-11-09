package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Set up the Application
	wms := app.New()
	window := wms.NewWindow("What's My Status?")
	window.CenterOnScreen()
	window.Resize(fyne.NewSize(800, 600))

	// Create the Input Fields
	statusMessageLabel := canvas.NewText("What's My Status?", color.White)
	statusMessageInput := canvas.NewText("Status", color.White)
	emojiLabel := canvas.NewText("What's my Emoji?", color.White)
	emojiInput := canvas.NewText("Emoji", color.White)
	// Top Container
	formContainer := container.New(layout.NewFormLayout(), statusMessageLabel, statusMessageInput, emojiLabel, emojiInput)
	topContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), formContainer, layout.NewSpacer())

	// Set status button
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
}
