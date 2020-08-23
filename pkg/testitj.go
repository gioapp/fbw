package main

import (
	"flag"
	"os"

	//"github.com/ViBiOh/auth/v2/pkg/ident/basic"
	//authMiddleware "github.com/ViBiOh/auth/v2/pkg/middleware"
	//basicMemory "github.com/ViBiOh/auth/v2/pkg/store/memory"
	"github.com/gioapp/fbw/pkg/crud"
	"github.com/gioapp/fbw/pkg/fibr"
	"github.com/gioapp/fbw/pkg/filesystem"
	"github.com/gioapp/fbw/pkg/renderer"
	"github.com/gioapp/fbw/pkg/thumbnail"
	//"github.com/gioapp/fbw/pkg/alcotest"
	//"github.com/gioapp/fbw/pkg/flags"
	//"github.com/gioapp/fbw/pkg/httputils"
	"github.com/gioapp/fbw/pkg/logger"
	//"github.com/gioapp/fbw/pkg/owasp"
	//"github.com/gioapp/fbw/pkg/prometheus"
)

//func newLoginApp(basicConfig basicMemory.Config) authMiddleware.App {
//	basicApp, err := basicMemory.New(basicConfig)
//	logger.Fatal(err)
//
//	basicProviderProvider := basic.New(basicApp)
//	return authMiddleware.New(basicApp, basicProviderProvider)
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

	fibrApp := files.New(crudApp, rendererApp, middlewareApp)

	go thumbnailApp.Start()
	go crudApp.Start()

	//server := httputils.New(serverConfig)
	//server.Middleware(prometheus.New(prometheusConfig).Middleware)
	//server.Middleware(owasp.New(owaspConfig).Middleware)
	//server.ListenServeWait(fibrApp.Handler())
}
