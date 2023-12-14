package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"os"
)

func main() {
	go func() {
		w := app.NewWindow(app.Title("PDF Merge"), app.Size(unit.Dp(800), unit.Dp(600)))
		var ops op.Ops
		var startButton widget.Clickable
		th := material.NewTheme()

		for {
			e := w.NextEvent()

			if e, ok := e.(system.FrameEvent); ok {
				gtx := layout.NewContext(&ops, e)
				btn := material.Button(th, &startButton, "Start")
				btn.Layout(gtx)
				e.Frame(gtx.Ops)
			}

			if _, ok := e.(system.DestroyEvent); ok {
				fmt.Println("Closing window")
				os.Exit(0)
			}
		}
	}()
	app.Main()
}
