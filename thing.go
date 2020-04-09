package main

import (
	"fmt"
	"gioui.org/f32"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"github.com/gioapp/gel"
	"github.com/gioapp/gelook"
	"image"
	"image/color"
	"time"
)

type DuoUIfbw struct {
	allSpaces   spaces
	details     *DuoUIfbwDetails
	theme       *gelook.DuoUItheme
	headerPath  string
	mainList    *layout.List
	topMenuList *layout.List
	detailsList *layout.List
	cursor      cursor
}

type cursor *DuoUIfbwThing

type DuoUIfbwThing struct {
	Name             string
	Type             string
	out              interface{}
	pressed          bool
	selected         bool
	details          *DuoUIfbwDetails
	check            *gel.CheckBox
	parentSpaceIndex int
	thingsFromSpace  []*DuoUIfbwThing
}

type DuoUIfbwDetails struct {
	filename string
	fullPath string
	ext      string
	Size     int64
	Mode     string
	ModTime  time.Time
}

func (f *DuoUIfbw) spaces() (s []DuoUIfbwSpace) {
	i := 0
	for index, space := range f.allSpaces {
		if index == i {
			s = append(s, space)
			i++
		}
	}
	return
}

func (t *DuoUIfbwThing) Layout(c *cursor, s *spaces, th *gelook.DuoUItheme, thingName string, col color.RGBA, gtx *layout.Context) {
	for _, e := range gtx.Events(t) {
		if e, ok := e.(pointer.Event); ok {
			switch e.Type {
			case pointer.Press:
				t.pressed = true
				t.selected = true
				*c = t
				if t.Type == "space" {
					s.NewDuoUIspace(t.parentSpaceIndex+1, "/"+t.Name)
					fmt.Println(t.details.fullPath)
					fmt.Println(t.details.filename)
					fmt.Println("oooooooooooooooooooooooooooooooooooooooooooooooo")
					fmt.Println(t.details)
				}

			case pointer.Release:
				t.pressed = false
			}
		}
	}
	if t.pressed {
		col = color.RGBA{A: 0xff, R: 0x30, G: 0xcf, B: 0xcf}

	}
	pointer.Rect(
		image.Rectangle{Max: image.Point{X: 500, Y: 500}},
	).Add(gtx.Ops)
	pointer.InputOp{Key: t}.Add(gtx.Ops)
	drawSquare(gtx.Ops, col)
	th.DuoUIcheckBox(thingName, "", "").Layout(gtx, t.check)
}

func drawSquare(ops *op.Ops, color color.RGBA) {
	square := f32.Rectangle{
		Max: f32.Point{X: 500, Y: 500},
	}
	paint.ColorOp{Color: color}.Add(ops)
	paint.PaintOp{Rect: square}.Add(ops)
}
