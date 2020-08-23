package files

import (
	"errors"
	"fmt"
	"github.com/ViBiOh/auth/v2/pkg/auth"

	"net/http"
	"strings"

	//authMiddleware "github.com/gioapp/fbw/pkg/middleware"
	"github.com/gioapp/fbw/pkg/crud"
	"github.com/gioapp/fbw/pkg/provider"
	"github.com/gioapp/fbw/pkg/renderer"
)

// App of package
type App interface {
	Handler(string) provider.Page
}

type app struct {
	//loginApp    authMiddleware.App
	crudApp     crud.App
	rendererApp renderer.App
}

// New creates new App from Config
func New(crudApp crud.App, rendererApp renderer.App) App {
	return &app{
		crudApp:     crudApp,
		rendererApp: rendererApp,
		//loginApp:    loginApp,
	}
}

func (a app) parseShare(request *provider.Request, authorizationHeader string) error {
	share := a.crudApp.GetShare(request.Path)
	if share == nil {
		return nil
	}

	if err := share.CheckPassword(authorizationHeader); err != nil {
		return err
	}

	request.Share = share
	request.CanEdit = share.Edit
	request.Path = strings.TrimPrefix(request.Path, fmt.Sprintf("/%s", share.ID))

	return nil
}

func convertAuthenticationError(err error) *provider.Error {
	if errors.Is(err, auth.ErrForbidden) {
		return provider.NewError(http.StatusForbidden, errors.New("you're not authorized to speak to me"))
	}

	//if errors.Is(err, ident.ErrMalformedAuth) {
	//	return provider.NewError(http.StatusBadRequest, err)
	//}

	return provider.NewError(http.StatusUnauthorized, err)
}

func (a app) parseRequest() (provider.Request, *provider.Error) {
	preferences := provider.Preferences{}
	//if cookie, err := r.Cookie("list_layout_paths"); err == nil {
	//	if value := cookie.Value; len(value) > 0 {
	//		preferences.ListLayoutPath = strings.Split(value, ",")
	//	}
	//}

	request := provider.Request{
		//Path:        r.URL.Path,
		CanEdit:  false,
		CanShare: false,
		//Display:     r.URL.Query().Get("d"),
		Preferences: preferences,
	}

	//if err := a.parseShare(&request, r.Header.Get("Authorization")); err != nil {
	//	return request, provider.NewError(http.StatusUnauthorized, err)
	//}

	if request.Share != nil {
		return request, nil
	}

	//if a.loginApp == nil {
	//	request.CanEdit = true
	//	request.CanShare = true
	//	return request, nil
	//}

	//_, user, err := a.loginApp.IsAuthenticated(r, "")
	//if err != nil {
	//	return request, convertAuthenticationError(err)
	//}

	//if a.loginApp.HasProfile(r.Context(), user, "admin") {
	//	request.CanEdit = true
	//	request.CanShare = true
	//}

	return request, nil
}

func (a app) handleRequest(request provider.Request) {
	//switch r.Method {
	//case http.MethodGet:
	//a.crudApp.Get(request)
	//case http.MethodPost:
	//	a.crudApp.Post(w, r, request)
	//case http.MethodPut:
	//	a.crudApp.Create(w, r, request)
	//case http.MethodPatch:
	//	a.crudApp.Rename(w, r, request)
	//case http.MethodDelete:
	//	a.crudApp.Delete(w, r, request)
	//default:
	//	httperror.NotFound(w)
	//}
}

// Handler for request. Should be use with net/http
func (a app) Handler(path string) provider.Page {
	//request, err := a.parseRequest()
	//if err != nil {
	//	a.rendererApp.Error(request, err)
	//	return
	//}
	preferences := provider.Preferences{}

	request := provider.Request{
		Path:     path,
		CanEdit:  false,
		CanShare: false,
		//Display:     r.URL.Query().Get("d"),
		Preferences: preferences,
	}

	a.handleRequest(request)

	return a.crudApp.Get(request)
}
