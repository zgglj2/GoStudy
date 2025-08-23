//go:build !wasm
// +build !wasm

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Fyne Test")

	label := widget.NewLabel("Hello Fyne!")
	button := widget.NewButton("Click Me", func() {
		label.SetText("Button Clicked!")
	})

	content := container.NewVBox(
		label,
		button,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 200))
	myWindow.ShowAndRun()
}
