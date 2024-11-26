package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func main() {
	a := app.New()
	w := a.NewWindow("Application")

	background := canvas.NewRectangle(color.NRGBA{R: 64, G: 58, B: 150, A: 58})
	backgroundImage := canvas.NewImageFromFile("maxresdefault.jpg")
	backgroundImage.FillMode = canvas.ImageFillOriginal
	label := widget.NewLabel("i see you")
	iconImage, _ := fyne.LoadResourceFromPath("unnamed.jpg")
	w.SetIcon(iconImage)
	w.SetContent(widget.NewLabel("xxxxxxxdddd"))
	w.Resize(fyne.NewSize(800, 600))
	//сделать кнопку
	//залить на гит
	//gray := color.RGBA{R: 200, G: 200, B: 200, A: 255}
	btn := widget.NewButton("Click me", func() {
		println("Button clicked")
		w.Close()
	})

	content := container.NewMax(
		btn,
		background,
		backgroundImage,
		container.NewCenter(label),
	)
	w.SetContent(content)
	w.ShowAndRun()
}
