package main

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/fbw/pkg/files"
	"time"
)

type fbw struct {
	fibr           files.App
	theme          *material.Theme
	iconFolder     *widget.Icon
	iconFolderOpen *widget.Icon
	iconFile       *widget.Icon
	iconMenu       *widget.Icon
	headerPath     string
	mainList       *layout.List
	topMenuList    *layout.List
	detailsList    *layout.List
	path           []string
}

type fbwThing struct {
	Name             string
	Type             string
	out              interface{}
	pressed          bool
	selected         bool
	details          *fbwDetails
	check            *widget.Bool
	parentSpaceIndex int
	thingsFromSpace  []*fbwThing
}

type fbwDetails struct {
	filename string
	fullPath string
	ext      string
	Size     int64
	Mode     string
	ModTime  time.Time
}
