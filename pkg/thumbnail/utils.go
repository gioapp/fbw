package thumbnail

import (
	"fmt"
	"path"
	"strings"

	"github.com/gioapp/fbw/pkg/provider"
)

// CanHaveThumbnail determine if thumbnail can be generated for given pathname
func CanHaveThumbnail(item provider.StorageItem) bool {
	return item.IsImage() || item.IsPdf() || item.IsVideo()
}

// HasThumbnail determine if thumbnail exist for given pathname
func (a app) HasThumbnail(item provider.StorageItem) bool {
	if !a.Enabled() {
		return false
	}

	_, err := a.storage.Info(getThumbnailPath(item))
	return err == nil
}

func getThumbnailPath(item provider.StorageItem) string {
	fullPath := path.Join(provider.MetadataDirectoryName, item.Pathname)
	if item.IsDir {
		return fullPath
	}

	return fmt.Sprintf("%s.jpg", strings.TrimSuffix(fullPath, path.Ext(fullPath)))
}
