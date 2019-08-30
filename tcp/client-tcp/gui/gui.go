package gui

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/widget"
	tcp "github.com/PedroCosta8/sistemas-distribuidos/client-tcp"
)

type App struct {
	searchEntry *widget.Entry
	image       *canvas.Image
	notFound    *canvas.Text
}

func NewApp() *App {
	return &App{
		notFound: canvas.NewText("", color.White),
	}
}

func (app *App) GetText() string {
	return app.searchEntry.Text
}

func (app *App) LoadImage() {
	text := app.GetText()
	image, err := tcp.SearchImage(text)
	defer image.Close()
	if err != nil {
		app.notFound.Text = "Image not found"
		app.image.File = ""
		canvas.Refresh(app.notFound)
		return
	}
	app.notFound.Text = ""
	app.image.File = image.Name()
	app.image.Resize(fyne.NewSize(800, 400))
	canvas.Refresh(app.image)
}

func Show(app fyne.App) {
	application := NewApp()
	window := app.NewWindow("Image searcher")
	window.CenterOnScreen()
	window.Resize(fyne.NewSize(800, 500))

	textField := widget.NewEntry()
	textField.FocusGained()
	application.searchEntry = textField

	searchButton := widget.NewButton("Search", func() {
		application.LoadImage()
	})
	searchButton.Style = widget.PrimaryButton

	application.image = &canvas.Image{FillMode: canvas.ImageFillContain}
	window.SetContent(widget.NewVBox(
		textField,
		searchButton,
		application.image,
		application.notFound,
	))
	window.ShowAndRun()
}
