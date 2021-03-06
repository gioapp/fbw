package crudtest

import (
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/gioapp/fbw/pkg/provider"
	"golang.org/x/crypto/bcrypt"
)

var (
	// PasswordLessShare instance
	PasswordLessShare = &provider.Share{
		ID:       "a1b2c3d4f5",
		Edit:     false,
		RootName: "public",
		File:     false,
		Path:     "/public",
	}

	passwordHash, _ = bcrypt.GenerateFromPassword([]byte("password"), 12)

	// PasswordShare instance
	PasswordShare = &provider.Share{
		ID:       "f5d4c3b2a1",
		Edit:     true,
		RootName: "private",
		File:     false,
		Path:     "/private",
		Password: string(passwordHash),
	}
)

// App for mocked calls
type App struct{}

// New creates new mocked instance
func New() App {
	return App{}
}

// Start mocked implementation
func (a App) Start() {
}

// Browser mocked implementation
func (a App) Browser(provider.Request, provider.StorageItem, *provider.Message) {
}

// ServeStatic mocked implementation
func (a App) ServeStatic(*http.Request) bool {
	return false
}

// List mocked implementation
func (a App) List(provider.Request, *provider.Message) {
}

// Get mocked implementation
func (a App) Get(provider.Request) {
}

// Post mocked implementation
func (a App) Post(*http.Request, provider.Request) {
}

// Create mocked implementation
func (a App) Create(*http.Request, provider.Request) {
}

// Upload mocked implementation
func (a App) Upload(*http.Request, provider.Request, *multipart.Part) {
}

// Rename mocked implementation
func (a App) Rename(*http.Request, provider.Request) {
}

// Delete mocked implementation
func (a App) Delete(*http.Request, provider.Request) {
}

// GetShare mocked implementation
func (a App) GetShare(path string) *provider.Share {
	if strings.HasPrefix(path, "/a1b2c3d4f5") {
		return PasswordLessShare
	}

	if strings.HasPrefix(path, "/f5d4c3b2a1") {
		return PasswordShare
	}

	return nil
}

// CreateShare mocked implementation
func (a App) CreateShare(*http.Request, provider.Request) {
}

// DeleteShare mocked implementation
func (a App) DeleteShare(*http.Request, provider.Request) {
}
