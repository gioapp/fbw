package filesystem

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gioapp/fbw/pkg/flags"
	"github.com/gioapp/fbw/pkg/logger"
	"github.com/gioapp/fbw/pkg/provider"
)

var (
	// ErrRelativePath occurs when path is relative (contains ".."")
	ErrRelativePath = errors.New("pathname contains relatives paths")
)

// Config of package
type Config struct {
	directory *string
}

type app struct {
	rootDirectory string
	rootDirname   string

	ignoreFn func(provider.StorageItem) bool
}

// Flags adds flags for configuring package
func Flags(fs *flag.FlagSet, prefix string) Config {
	return Config{
		directory: flags.New(prefix, "filesystem").Name("Directory").Default("./").Label("Path to served directory").ToString(fs),
	}
}

// New creates new App from Config
func New(config Config) (provider.Storage, error) {
	rootDirectory := strings.TrimSpace(*config.directory)

	if len(rootDirectory) == 0 {
		return nil, errors.New("no directory provided")
	}

	info, err := os.Stat(rootDirectory)
	if err != nil {
		return nil, convertError(err)
	}

	if !info.IsDir() {
		return nil, fmt.Errorf("path %s is not a directory", rootDirectory)
	}

	logger.Info("Serving file from %s", rootDirectory)

	return &app{
		rootDirectory: rootDirectory,
		rootDirname:   info.Name(),
	}, nil
}

func (a *app) SetIgnoreFn(ignoreFn func(provider.StorageItem) bool) {
	a.ignoreFn = ignoreFn
}

// Info provide metadata about given pathname
func (a *app) Info(pathname string) (provider.StorageItem, error) {
	if err := a.checkPathname(pathname); err != nil {
		return provider.StorageItem{}, convertError(err)
	}

	fullpath := a.getFullPath(pathname)

	info, err := os.Stat(fullpath)
	if err != nil {
		return provider.StorageItem{}, convertError(err)
	}

	return convertToItem(a.getRelativePath(fullpath), info), nil
}

// List items in the storage
func (a *app) List(pathname string) ([]provider.StorageItem, error) {
	if err := a.checkPathname(pathname); err != nil {
		return nil, convertError(err)
	}

	fullpath := a.getFullPath(pathname)

	files, err := ioutil.ReadDir(fullpath)
	if err != nil {
		return nil, convertError(err)
	}

	items := make([]provider.StorageItem, 0)
	for _, file := range files {
		item := convertToItem(a.getRelativePath(path.Join(fullpath, file.Name())), file)
		if a.ignoreFn != nil && a.ignoreFn(item) {
			continue
		}

		items = append(items, item)
	}

	sort.Sort(ByHybridSort(items))

	return items, nil
}

// WriterTo opens writer for given pathname
func (a *app) WriterTo(pathname string) (io.WriteCloser, error) {
	if err := a.checkPathname(pathname); err != nil {
		return nil, convertError(err)
	}

	writer, err := a.getFile(pathname)
	return writer, convertError(err)
}

// ReaderFrom reads content from given pathname
func (a *app) ReaderFrom(pathname string) (provider.ReadSeekerCloser, error) {
	if err := a.checkPathname(pathname); err != nil {
		return nil, convertError(err)
	}

	output, err := os.OpenFile(a.getFullPath(pathname), os.O_RDONLY, getMode(pathname))
	return output, convertError(err)
}

// Walk browses item recursively
func (a *app) Walk(pathname string, walkFn func(provider.StorageItem, error) error) error {
	pathname = path.Join(a.rootDirectory, pathname)

	return convertError(filepath.Walk(pathname, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logger.Error("%s", err)
			return walkFn(provider.StorageItem{}, err)
		}

		item := convertToItem(a.getRelativePath(path), info)
		if a.ignoreFn != nil && a.ignoreFn(item) {
			if item.IsDir {
				return filepath.SkipDir
			}
			return nil
		}

		return walkFn(item, err)
	}))
}

// Create container in storage
func (a *app) CreateDir(name string) error {
	if err := a.checkPathname(name); err != nil {
		return convertError(err)
	}

	return convertError(os.MkdirAll(a.getFullPath(name), 0700))
}

// Store file to storage
func (a *app) Store(pathname string, content io.ReadCloser) error {
	if err := a.checkPathname(pathname); err != nil {
		return convertError(err)
	}

	storageFile, err := a.getFile(pathname)
	if storageFile != nil {
		defer func() {
			if err := storageFile.Close(); err != nil {
				logger.Error("unable to close stored file: %s", err)
			}
		}()
	}

	if err != nil {
		return convertError(err)
	}

	copyBuffer := make([]byte, 32*1024)
	if _, err = io.CopyBuffer(storageFile, content, copyBuffer); err != nil {
		return convertError(err)
	}

	return nil
}

// Rename file or directory from storage
func (a *app) Rename(oldName, newName string) error {
	if err := a.checkPathname(oldName); err != nil {
		return convertError(err)
	}

	if err := a.checkPathname(newName); err != nil {
		return convertError(err)
	}

	return convertError(os.Rename(a.getFullPath(oldName), a.getFullPath(newName)))
}

// Remove file or directory from storage
func (a *app) Remove(pathname string) error {
	if err := a.checkPathname(pathname); err != nil {
		return convertError(err)
	}

	return convertError(os.RemoveAll(a.getFullPath(pathname)))
}
