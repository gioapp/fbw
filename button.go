package main

import (
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/gioapp/gelook"
	"image"
)

type Button struct {
	pressed        bool
	Name           string
	Do             func(interface{})
	ColorBg        string
	ColorBgHover   string
	ColorText      string
	ColorTextHover string
	BorderRadius   [4]float32
	OperateValue   interface{}
}

func (b *Button) Layout(gtx *layout.Context, th *gelook.DuoUItheme) {
	for _, e := range gtx.Events(b) { // HLevent
		if e, ok := e.(pointer.Event); ok { // HLevent
			switch e.Type { // HLevent
			case pointer.Press: // HLevent
				b.pressed = true // HLevent
				//b.Do(b.OperateValue)
			case pointer.Release: // HLevent
				b.pressed = false // HLevent
			}
		}
	}

	cs := gtx.Constraints

	//colorBg := helpers.HexARGB("ff30cfcf")
	//colorText := HexARGB(b.ColorText)
	colorBg := HexARGB(b.ColorBg)

	if b.pressed {
		//colorText = HexARGB(b.ColorTextHover)
		colorBg = HexARGB(b.ColorBgHover)
	}
	pointer.Rect( // HLevent
		image.Rectangle{Max: image.Point{X: cs.Width.Max, Y: cs.Height.Max}}, // HLevent
	).Add(gtx.Ops) // HLevent
	pointer.InputOp{Key: b}.Add(gtx.Ops) // HLevent
	DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, colorBg, b.BorderRadius, unit.Dp(0))

	in := layout.UniformInset(unit.Dp(5))
	in.Layout(gtx, func() {
		DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, colorBg, b.BorderRadius, unit.Dp(0))
		//cs := gtx.Constraints
		txt := th.Caption(b.Name)
		txt.Color = "ff303030"
		txt.Layout(gtx)
	})
}
