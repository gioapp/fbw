package main

import (
	"gioui.org/app"
	_ "gioui.org/app/permission/storage"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"golang.org/x/exp/shiny/materialdesign/icons"

	"log"
)

func Newfbw(path string) *fbw {
	return &fbw{
		//allSpaces: make(map[int]fbwSpace),
		theme:          material.NewTheme(gofont.Collection()),
		iconFolder:     mustIcon(widget.NewIcon(icons.FileFolder)),
		iconFolderOpen: mustIcon(widget.NewIcon(icons.FileFolderOpen)),
		iconFile:       mustIcon(widget.NewIcon(icons.FileAttachment)),

		topMenuList: &layout.List{
			Axis: layout.Horizontal,
		},
		mainList: &layout.List{
			Axis: layout.Horizontal,
		},
		detailsList: &layout.List{
			Axis: layout.Horizontal,
		},
		path: []string{path},
	}
}

//col := color.RGBA{A: 0xff, R: 0x30, G: 0xcf, B: 0xcf}
//fbw.spaces = append(fbw.spaces, root)
//fbw.headerPath = fbw.cursor.Name
//itemValue := item{
//	i: 0,
//}
//list := &layout.List{
//	Axis: layout.Horizontal,
//}
func main() {
	//currentFolder = make(chan I)
	f := Newfbw("")
	//fbw.allSpaces.NewDuoUIspace(0, "/")
	//root := newfbwSpace("/")
	//root.getThingsFromSpace()

	makeList(getThingsFromSpace(f.path))
	go func() {
		if err := f.loop(app.NewWindow(app.Size(unit.Dp(800), unit.Dp(650)))); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

func (f *fbw) loop(w *app.Window) error {
	var o op.Ops
	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				c := layout.NewContext(&o, e)
				f.appMain(c)
				e.Frame(c.Ops)
			}
		}
	}
}

func mustIcon(ic *widget.Icon, err error) *widget.Icon {
	if err != nil {
		panic(err)
	}
	return ic
}
