package main

import (
	"gioui.org/layout"
	"github.com/gioapp/gel"
	"github.com/gioapp/gelook"
	"image/color"
	"io/ioutil"
	"log"
	"os"
)

type spaces map[int]DuoUIfbwSpace

func (s spaces) NewDuoUIspace(i int, path string) {
	space := openDuoUIfbwSpace(i, path)
	space.getThingsFromSpace()
	s[i] = space
}

type DuoUIfbwSpace struct {
	displayHorizontalListIndex int
	path                       string
	parent                     string
	fullPath                   string
	thingsFromSpace            map[string]*DuoUIfbwThing
	listThingsFromSpace        []os.FileInfo
	list                       *layout.List
	child                      *DuoUIfbwSpace
}

func openDuoUIfbwSpace(i int, path string) DuoUIfbwSpace {
	return DuoUIfbwSpace{
		displayHorizontalListIndex: i,
		path:                       path,
		thingsFromSpace:            make(map[string]*DuoUIfbwThing),
		list: &layout.List{
			Axis: layout.Vertical,
		},
		child: &DuoUIfbwSpace{},
	}
}

func (s *DuoUIfbwSpace) listThings(c *cursor, ss *spaces, gtx *layout.Context, th *gelook.DuoUItheme) func() {
	return func() {
		s.list.Layout(gtx, len(s.listThingsFromSpace), func(i int) {
			col := color.RGBA{A: 0xff, R: 0xcf, G: 0xcf, B: 0x30}
			if s.listThingsFromSpace[i].IsDir() {
			} else {
				col = color.RGBA{A: 0xff, R: 0xcf, G: 0x30, B: 0x30}
			}
			s.thingsFromSpace[s.listThingsFromSpace[i].Name()].Layout(c, ss, th, s.listThingsFromSpace[i].Name(), col, gtx)
		})
	}
}

func (s *DuoUIfbwSpace) getThingsFromSpace() (err error) {
	if s.path == "" {
		s.path = ""
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
			Name:             t.Name(),
			details:          details,
			check:            new(gel.CheckBox),
			parentSpaceIndex: s.displayHorizontalListIndex,
		}
		if t.IsDir() {
			s.thingsFromSpace[t.Name()].Type = "space"
		}
	}
	return
}
