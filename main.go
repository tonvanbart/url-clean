package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/atotto/clipboard"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("URL cleaner")

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter URL here")

	result := widget.NewLabel("")

	button := widget.NewButton("Clean & copy", func() {
		transformed, _, _ := strings.Cut(input.Text, "?")
		result.SetText(transformed)

		// Copy to clipboard
		clipboard.WriteAll(transformed)
	})
	button.Importance = widget.HighImportance

	myWindow.SetContent(container.NewVBox(
		widget.NewLabel("Enter a URL to clean"),
		input,
		button,
		result,
	))

	myWindow.Resize(fyne.NewSize(500, 200))
	myWindow.ShowAndRun()
}
