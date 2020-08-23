package crud

import (
	"net/http"
	"path"
	"strings"

	"github.com/gioapp/fbw/pkg/provider"
)

var (
	staticRootPath = []string{
		"/robots.txt",
		"/browserconfig.xml",
		"/favicon.ico",
	}
)

// ServeStatic check if filename match SEO or static filename and serve it
func (a *app) ServeStatic(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodGet {
		return false
	}

	if r.URL.Path == "/sitemap.xml" {
		a.renderer.Sitemap(w)
		return true
	}

	if strings.HasPrefix(r.URL.Path, "/svg") {
		a.renderer.SVG(strings.TrimPrefix(r.URL.Path, "/svg/"), r.URL.Query().Get("fill"))
		return true
	}

	if strings.HasPrefix(r.URL.Path, "/favicon") {
		http.ServeFile(w, r, path.Join("templates/static", r.URL.Path))
		return true
	}

	for _, staticPath := range staticRootPath {
		if r.URL.Path == staticPath {
			http.ServeFile(w, r, path.Join("templates/static", r.URL.Path))
			return true
		}
	}

	return false
}

func (a *app) getWithMessage(request provider.Request, message *provider.Message) provider.Page {
	//info, err := a.storage.Info(request.GetFilepath(""))
	//if err != nil {
	//	if provider.IsNotExist(err) {
	//		a.renderer.Error(request, provider.NewError(http.StatusNotFound, err))
	//	} else {
	//		a.renderer.Error(request, provider.NewError(http.StatusInternalServerError, err))
	//	}
	//	return
	//}

	//if query.GetBool( "thumbnail") {
	//	if info.IsDir {
	//		a.thumbnail.List( info)
	//	} else {
	//		a.thumbnail.Serve( info)
	//	}
	//
	//	return
	//}

	//if !info.IsDir {
	//	if query.GetBool( "browser") {
	//		a.Browser(request, info, message)
	//	//} else if file, err := a.storage.ReaderFrom(info.Pathname); err != nil {
	//	//	a.renderer.Error( request, provider.NewError(http.StatusInternalServerError, err))
	//	//} else {
	//		//http.ServeContent(w, r, info.Name, info.Date, file)
	//	}
	//
	//	return
	//}
	//
	//if query.GetBool( "download") {
	//	a.Download( request)
	//	return
	//}

	//if !strings.HasSuffix("", "/") {
	//	//http.Redirect( fmt.Sprintf("%s/", r.URL.Path), http.StatusPermanentRedirect)
	//	return
	//}

	return a.List(request, message)
}

// Get output content
func (a *app) Get(request provider.Request) provider.Page {
	var message *provider.Message

	if messageContent := strings.TrimSpace("/"); messageContent != "" {
		message = &provider.Message{
			//Level:   r.URL.Query().Get("messageLevel"),
			Content: messageContent,
		}
	}

	return a.getWithMessage(request, message)

	//fmt.Println(":ssssssssssssss:", ss)
	//
}
