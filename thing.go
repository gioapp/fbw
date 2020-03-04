package main

import (
	"gioui.org/f32"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"github.com/gioapp/fbw/theme"
	"image"
	"image/color"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type DuoUIfbw struct {
	spaces  []*DuoUIfbwSpace
	details *DuoUIfbwDetails
	theme   *theme.Theme
}

type DuoUIfbwSpace struct {
	path                string
	parent              string
	fullPath            string
	thingsFromSpace     map[string]*DuoUIfbwThing
	listThingsFromSpace []os.FileInfo
}

type DuoUIfbwThing struct {
	Name            string
	Type            string
	out             interface{}
	pressed         bool
	selected        bool
	details         *DuoUIfbwDetails
	check           *widget.CheckBox
	thingsFromSpace []*DuoUIfbwThing
}
type DuoUIfbwDetails struct {
	filename string
	fullPath string
	ext      string
	Size     int64
	Mode     string
	ModTime  time.Time
}

func (t *DuoUIfbwThing) Layout(th *theme.Theme, thingName string, col color.RGBA, gtx *layout.Context) {
	for _, e := range gtx.Events(t) {
		if e, ok := e.(pointer.Event); ok {
			switch e.Type {
			case pointer.Press:
				t.pressed = true
				t.selected = true
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
	th.CheckBox(thingName).Layout(gtx, t.check)
}

func drawSquare(ops *op.Ops, color color.RGBA) {
	square := f32.Rectangle{
		Max: f32.Point{X: 500, Y: 500},
	}
	paint.ColorOp{Color: color}.Add(ops)
	paint.PaintOp{Rect: square}.Add(ops)
}

func (s *DuoUIfbwSpace) listThings(gtx *layout.Context, th *theme.Theme) func() {
	return func() {
		list := &layout.List{
			Axis: layout.Vertical,
		}
		list.Layout(gtx, len(s.listThingsFromSpace), func(i int) {
			col := color.RGBA{A: 0xff, R: 0xcf, G: 0xcf, B: 0x30}
			if s.listThingsFromSpace[i].IsDir() {
			} else {
				col = color.RGBA{A: 0xff, R: 0xcf, G: 0x30, B: 0x30}
			}
			s.thingsFromSpace[s.listThingsFromSpace[i].Name()].Layout(th, s.listThingsFromSpace[i].Name(), col, gtx)
		})
	}
}

func (s *DuoUIfbwSpace) getThingsFromSpace() (err error) {
	if s.path == "" {
		s.path = "/"
	}
	s.listThingsFromSpace, err = ioutil.ReadDir(s.path)
	if err != nil {
		log.Fatal(err)
	}
	s.thingsFromSpace = make(map[string]*DuoUIfbwThing)
	for _, t := range s.listThingsFromSpace {
		details := &DuoUIfbwDetails{
			Mode:    t.Mode().String(),
			Size:    t.Size(),
			ModTime: t.ModTime(),
		}
		s.thingsFromSpace[t.Name()] = &DuoUIfbwThing{
			Name:    t.Name(),
			details: details,
			check:   new(widget.CheckBox),
		}
		if t.IsDir() {
			s.thingsFromSpace[t.Name()].Type = "space"
		}
	}
	return
}
