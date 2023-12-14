package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/unit"
	"os"
)

func main() {
	go func() {
		w := app.NewWindow(app.Title("PDF Merge"), app.Size(unit.Dp(800), unit.Dp(600)))
		for {
			e := w.NextEvent()
			switch e.(type) {
			case system.DestroyEvent:
				fmt.Println("Closing window")
				os.Exit(0)
			}
		}
	}()
	app.Main()
}
