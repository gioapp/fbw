package main

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/widget"
	"github.com/gioapp/fbw/pkg/item"
	"github.com/gioapp/fbw/pkg/provider"
	"github.com/gioapp/gel/helper"
	"github.com/w-ingsolutions/c/pkg/lyt"
	"os"
	"time"
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
	Date  time.Time
	Dir   bool
	Type  string
	Btn   *widget.Clickable
	Check *widget.Bool
}

func (f *fbw) appMain(gtx layout.Context) layout.Dimensions {
	//fmt.Println("kkk", items)
	return lyt.Format(gtx, "vflexb(middle,r(_),f(1,_))",
		func(gtx layout.Context) layout.Dimensions {
			return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
				func(gtx layout.Context) layout.Dimensions {
					b := item.ItemBtn(f.theme, upBtn, checkAll, f.iconFile, f.iconFolder, "", "..", 0).Layout(gtx)
					for upBtn.Clicked() {
						//parentCid = c
						f.back()
						//fmt.Println("Path", f.path)
						f.handle()
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
						b := item.ItemBtn(f.theme, itm.Btn, itm.Check, f.iconFile, f.iconFolder, itm.Name, itm.Type, itm.Size).Layout(gtx)
						for itm.Btn.Clicked() {
							f.path = append(f.path, itm.Name)
							//makeList(getThingsFromSpace(f.path))
							f.handle()

						}
						return b
					},
					helper.DuoUIline(false, 0, 0, 1, "ff995577"),
				)
			})
		})
	return layout.Dimensions{}

}
func (f *fbw) handle() {
	var it I
	for _, item := range f.fibr.Handler(path(f.path)).Content["Files"].([]provider.RenderItem) {
		if string(item.Name[0]) != "." {
			it = append(it, &FolderListItem{
				Name:  item.Name,
				Date:  item.Date,
				Dir:   item.IsDir,
				Type:  item.Mime(),
				Size:  item.Info.(os.FileInfo).Size(),
				Btn:   new(widget.Clickable),
				Check: new(widget.Bool),
			})
		}
	}
	currentFolder = &it
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
}

func (f *fbw) back() {

	if len(f.path) > 0 {
		f.path = f.path[:len(f.path)-1]
	}

	fmt.Println("DDD", f.path)
	return
}

func path(pathNames []string) string {
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
	return path
}
