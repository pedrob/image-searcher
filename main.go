package main

import (
	"fyne.io/fyne/app"
	"github.com/PedroCosta8/sistemas-distribuidos/client-tcp/gui"
)

func main() {
	app := app.New()
	gui.Show(app)
}
