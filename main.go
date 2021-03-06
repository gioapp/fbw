package main

import (
	"flag"
	"gioui.org/app"
	_ "gioui.org/app/permission/storage"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/fbw/pkg/logger"
	"golang.org/x/exp/shiny/materialdesign/icons"
	"log"
	"os"
	//"github.com/ViBiOh/auth/v2/pkg/ident/basic"
	//authMiddleware "github.com/ViBiOh/auth/v2/pkg/middleware"
	//basicMemory "github.com/ViBiOh/auth/v2/pkg/store/memory"
	"github.com/gioapp/fbw/pkg/crud"
	"github.com/gioapp/fbw/pkg/files"
	"github.com/gioapp/fbw/pkg/filesystem"
	"github.com/gioapp/fbw/pkg/renderer"
	"github.com/gioapp/fbw/pkg/thumbnail"
)

//func newLoginApp(basicConfig basicMemory.Config) authMiddleware.App {
//	basicApp, err := basicMemory.New(basicConfig)
//	logger.Fatal(err)
//
//	basicProviderProvider := basic.New(basicApp)
//	return authMiddleware.New(basicApp, basicProviderProvider)
//}

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

	fs := flag.NewFlagSet("fibr", flag.ExitOnError)

	//serverConfig := httputils.Flags(fs, "")
	//alcotestConfig := alcotest.Flags(fs, "")
	//prometheusConfig := prometheus.Flags(fs, "prometheus")
	//owaspConfig := owasp.Flags(fs, "")

	//basicConfig := basicMemory.Flags(fs, "auth")

	crudConfig := crud.Flags(fs, "")
	rendererConfig := renderer.Flags(fs, "")

	filesystemConfig := filesystem.Flags(fs, "fs")
	thumbnailConfig := thumbnail.Flags(fs, "thumbnail")

	//disableAuth := flags.New("", "auth").Name("NoAuth").Default(false).Label("Disable basic authentification").ToBool(fs)

	logger.Fatal(fs.Parse(os.Args[1:]))

	//alcotest.DoAndExit(alcotestConfig)

	storage, err := filesystem.New(filesystemConfig)
	logger.Fatal(err)

	thumbnailApp := thumbnail.New(thumbnailConfig, storage)
	rendererApp := renderer.New(rendererConfig, thumbnailApp)
	crudApp, err := crud.New(crudConfig, storage, rendererApp, thumbnailApp)
	logger.Fatal(err)

	//var middlewareApp authMiddleware.App
	//if !*disableAuth {
	//	middlewareApp = newLoginApp(basicConfig)
	//}
	f := Newfbw("")

	f.fibr = files.New(crudApp, rendererApp)

	go thumbnailApp.Start()
	go crudApp.Start()

	//server := httputils.New(serverConfig)
	//server.Middleware(prometheus.New(prometheusConfig).Middleware)
	//server.Middleware(owasp.New(owaspConfig).Middleware)
	//server.ListenServeWait(fibrApp.Handler())

	//currentFolder = make(chan I)

	//fbw.allSpaces.NewDuoUIspace(0, "/")
	//root := newfbwSpace("/")
	//root.getThingsFromSpace()

	f.handle()

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
