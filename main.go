package main

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/gioapp/fbw/theme"
)

func main() {
	gofont.Register()

	fbw := &DuoUIfbw{
		theme: theme.NewTheme(),
	}
	root := &DuoUIfbwSpace{
		path:            "/",
		parent:          "/",
		fullPath:        "/",
		thingsFromSpace: make(map[string]*DuoUIfbwThing),
	}
	root.getThingsFromSpace()

	fbw.spaces = append(fbw.spaces, root)
	go func() {
		w := app.NewWindow(
			app.Size(unit.Dp(400), unit.Dp(800)),
			app.Title("ParallelCoin"),
		)

		mainList := &layout.List{
			Axis: layout.Horizontal,
		}
		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)

				widgets := []func(){
					fbw.spaces[0].listThings(gtx, fbw.theme),
					func() {
						//selected.Info(selected, col, gtx)
					},
				}
				mainList.Layout(gtx, len(widgets), func(i int) {
					layout.UniformInset(unit.Dp(0)).Layout(gtx, widgets[i])
				})
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
