package crud

import (
	"crypto/rand"
	"fmt"
	"github.com/gioapp/fbw/pkg/provider"
	"github.com/gioapp/fbw/pkg/sha"
	"path"
)

func uuid() (string, error) {
	raw := make([]byte, 16)
	_, _ = rand.Read(raw)

	raw[8] = raw[8]&^0xc0 | 0x80
	raw[6] = raw[6]&^0xf0 | 0x40

	return fmt.Sprintf("%x-%x-%x-%x-%x", raw[0:4], raw[4:6], raw[6:8], raw[8:10], raw[10:]), nil
}

// CreateShare create a share for given URL
func (a *app) CreateShare(request provider.Request) {
	if !request.CanShare {
		//a.renderer.Error(w, request, provider.NewError(http.StatusForbidden, ErrNotAuthorized))
		return
	}

	var err error

	edit := false
	//if editValue := strings.TrimSpace(r.FormValue("edit")); editValue != "" {
	//	edit, err = strconv.ParseBool(editValue)
	//	if err != nil {
	//a.renderer.Error(w, request, provider.NewError(http.StatusBadRequest, err))
	//return
	//}
	//}

	password := ""
	//if passwordValue := strings.TrimSpace(r.FormValue("password")); passwordValue != "" {
	//	hash, err := bcrypt.GenerateFromPassword([]byte(passwordValue), 12)
	//	if err != nil {
	//		//a.renderer.Error(w, request, provider.NewError(http.StatusInternalServerError, err))
	//		return
	//	}
	//
	//	password = string(hash)
	//}

	uuid, err := uuid()
	if err != nil {
		//a.renderer.Error(w, request, provider.NewError(http.StatusInternalServerError, err))
		return
	}
	id := sha.Sha1(uuid)[:8]

	a.metadataLock.Lock()
	defer a.metadataLock.Unlock()

	info, err := a.storage.Info(request.Path)
	if err != nil {
		if provider.IsNotExist(err) {
			//a.renderer.Error(w, request, provider.NewError(http.StatusNotFound, err))
		} else {
			//a.renderer.Error(w, request, provider.NewError(http.StatusInternalServerError, err))
		}
		return
	}

	a.metadatas = append(a.metadatas, &provider.Share{
		ID:       id,
		Path:     request.Path,
		RootName: path.Base(request.Path),
		Edit:     edit,
		Password: password,
		File:     !info.IsDir,
	})

	if err = a.saveMetadata(); err != nil {
		//a.renderer.Error(w, request, provider.NewError(http.StatusInternalServerError, err))
		return
	}

	//http.Redirect(w, r, fmt.Sprintf("%s/?message=%s&messageLevel=success#share-list", path.Dir(request.GetURI("")), url.QueryEscape(fmt.Sprintf("Share successfully created with ID: %s", id))), http.StatusFound)
}

// DeleteShare delete a share from given ID
func (a *app) DeleteShare(request provider.Request) {
	if !request.CanShare {
		//a.renderer.Error(w, request, provider.NewError(http.StatusForbidden, ErrNotAuthorized))
		return
	}

	//id := r.FormValue("id")

	a.metadataLock.Lock()
	defer a.metadataLock.Unlock()

	//for i, metadata := range a.metadatas {
	//if metadata.ID == id {
	//	a.metadatas = append(a.metadatas[:i], a.metadatas[i+1:]...)
	//	break
	//}
	//}

	if err := a.saveMetadata(); err != nil {
		//a.renderer.Error(w, request, provider.NewError(http.StatusInternalServerError, err))
		return
	}

	//http.Redirect(w, r, fmt.Sprintf("%s/?message=%s&messageLevel=success#share-list", request.GetURI(""), url.QueryEscape(fmt.Sprintf("Share with id %s successfully deleted", id))), http.StatusFound)
}
