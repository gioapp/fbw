package crud

import (
	"github.com/gioapp/fbw/pkg/provider"
)

// Browser render file web view
func (a *app) Browser(request provider.Request, file provider.StorageItem, message *provider.Message) {
	var (
	//previous *provider.StorageItem
	//next     *provider.StorageItem
	)

	//pathParts := getPathParts(request.GetURI(""))
	//breadcrumbs := pathParts[:len(pathParts)-1]

	//files, err := a.storage.List(path.Dir(file.Pathname))
	//if err != nil {
	//	logger.Error("unable to list neighbors files: %s", err)
	//} else {
	//previous, next = getPreviousAndNext(file, files)
	//}

	//content := map[string]interface{}{
	//	"Paths": breadcrumbs,
	//	"File": provider.RenderItem{
	//		ID:          sha.Sha1(file.Name),
	//		StorageItem: file,
	//	},
	//	"Cover":    a.getCover(files),
	//	"Parent":   path.Join(breadcrumbs...),
	//	"Previous": previous,
	//	"Next":     next,
	//}

	//a.renderer.File(w, request, content, message)
}
