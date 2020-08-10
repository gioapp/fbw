package main

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/widget"
	"github.com/gioapp/gel/helper"
	"github.com/w-ingsolutions/c/pkg/lyt"
	"io/ioutil"
	"os"
)

var (
	l = &layout.List{
		Axis: layout.Vertical,
	}
	upBtn    = new(widget.Clickable)
	checkAll = new(widget.Bool)

	currentFolder *I
)

type I []*FolderListItem

type FolderListItem struct {
	Name  string
	Size  int64
	Dir   bool
	Type  uint8
	Btn   *widget.Clickable
	Check *widget.Bool
}

func getThingsFromSpace(pathNames []string) []os.FileInfo {
	//if path == "" {
	//	path = "/"
	//}

	path := ""
	if len(pathNames) > 0 {
		fmt.Println("PaNNNNNDaaa")
		for _, pathName := range pathNames {
			path = path + string(os.PathSeparator) + pathName
		}
	} else {
		fmt.Println("PaDaaa")
		path = string(os.PathSeparator)
	}

	list, err := ioutil.ReadDir(path)
	checkError(err)
	//s.thingsFromSpace = make(map[string]*fbwThing)
	//for _, t := range s.listThingsFromSpace {
	//	details := &fbwDetails{
	//		Mode:    t.Mode().String(),
	//		Size:    t.Size(),
	//		ModTime: t.ModTime(),
	//	}
	//	s.thingsFromSpace[t.Name()] = &fbwThing{
	//		Name:    t.Name(),
	//		details: details,
	//		check:            new(widgetBool),
	//parentSpaceIndex: s.displayHorizontalListIndex,
	//}
	//if t.IsDir() {
	//	s.thingsFromSpace[t.Name()].Type = "space"
	//}
	//}
	return list
}

func makeList(items []os.FileInfo) {
	var it I
	for _, item := range items {
		if string(item.Name()[0]) != "." {
			it = append(it, &FolderListItem{
				Name: item.Name(),
				Size: item.Size(),
				Dir:  item.IsDir(),
				//Type:  uint8,

				Btn:   new(widget.Clickable),
				Check: new(widget.Bool),
			})
		}
	}
	currentFolder = &it

}

func (f *fbw) appMain(gtx layout.Context) layout.Dimensions {
	//fmt.Println("kkk", items)
	return lyt.Format(gtx, "vflexb(middle,r(_),f(1,_))",
		func(gtx layout.Context) layout.Dimensions {
			return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
				func(gtx layout.Context) layout.Dimensions {
					b := ItemBtn(f.theme, upBtn, checkAll, f.iconFile, f.iconFolder, "..", 0, 0).Layout(gtx)
					for upBtn.Clicked() {
						//parentCid = c
						f.back()
						fmt.Println("Path", f.path)
						makeList(getThingsFromSpace(f.path))
					}
					return b
				},
				helper.DuoUIline(false, 0, 0, 1, "ff995577"),
			)
			return layout.Dimensions{}
		},
		func(gtx layout.Context) layout.Dimensions {
			cf := *currentFolder
			return l.Layout(gtx, len(cf), func(gtx layout.Context, i int) layout.Dimensions {
				itm := cf[i]
				return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
					func(gtx layout.Context) layout.Dimensions {
						b := ItemBtn(f.theme, itm.Btn, itm.Check, f.iconFile, f.iconFolder, itm.Name, itm.Type, itm.Size).Layout(gtx)
						for itm.Btn.Clicked() {
							f.path = append(f.path, itm.Name)
							makeList(getThingsFromSpace(f.path))
						}
						return b
					},
					helper.DuoUIline(false, 0, 0, 1, "ff995577"),
				)
			})
		})
	return layout.Dimensions{}

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
}

///////////////////////

//type spaces map[int]fbwSpace
//
//func (s spaces) NewDuoUIspace(i int, path string) {
//	space := openfbwSpace(i, path)
//	space.getThingsFromSpace()
//	s[i] = space
//}
//
//type fbwSpace struct {
//	displayHorizontalListIndex int
//	path                       string
//	parent                     string
//	fullPath                   string
//	thingsFromSpace            map[string]*fbwThing
//	listThingsFromSpace        []os.FileInfo
//	list                       *layout.List
//	child                      *fbwSpace
//}
//
//func openfbwSpace(i int, path string) fbwSpace {
//	return fbwSpace{
//		displayHorizontalListIndex: i,
//		path:                       path,
//		thingsFromSpace:            make(map[string]*fbwThing),
//		list: &layout.List{
//			Axis: layout.Vertical,
//		},
//		child: &fbwSpace{},
//	}
//}

//func (s *fbwSpace) listThings(c *cursor, ss *spaces, gtx layout.Context, th material.Theme) func() {
//	return func() {
//s.list.Layout(gtx, len(s.listThingsFromSpace), func(i int) {
//	col := color.RGBA{A: 0xff, R: 0xcf, G: 0xcf, B: 0x30}
//	if s.listThingsFromSpace[i].IsDir() {
//	} else {
//		col = color.RGBA{A: 0xff, R: 0xcf, G: 0x30, B: 0x30}
//	}
//s.thingsFromSpace[s.listThingsFromSpace[i].Name()].Layout(c, ss, th, s.listThingsFromSpace[i].Name(), col, gtx)
//})
//}/
//}
//
//func (s *fbwSpace) getThingsFromSpace() (err error) {
//	if s.path == "" {
//		s.path = ""
//	}
//	s.listThingsFromSpace, err = ioutil.ReadDir(s.path)
//	if err != nil {
//		log.Fatal(err)
//	}
//	s.thingsFromSpace = make(map[string]*fbwThing)
//	for _, t := range s.listThingsFromSpace {
//		details := &fbwDetails{
//			Mode:    t.Mode().String(),
//			Size:    t.Size(),
//			ModTime: t.ModTime(),
//		}
//		s.thingsFromSpace[t.Name()] = &fbwThing{
//			Name:    t.Name(),
//			details: details,
//			//check:            new(widgetBool),
//			parentSpaceIndex: s.displayHorizontalListIndex,
//		}
//		if t.IsDir() {
//			s.thingsFromSpace[t.Name()].Type = "space"
//		}
//	}
//	return
//}

func (f *fbw) back() {

	if len(f.path) > 0 {
		f.path = f.path[:len(f.path)-1]
	}

	fmt.Println("DDD", f.path)
	return
}
