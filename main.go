package main

import (
	"gioui.org/app"
)

func main() {
	go func() {
		w := app.NewWindow()

		for {
			w.NextEvent()
		}
	}()
	app.Main()
}
