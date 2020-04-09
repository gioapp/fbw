package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/gioapp/gelook"
)

func NewDuoUIfbw() *DuoUIfbw {
	return &DuoUIfbw{
		allSpaces: make(map[int]DuoUIfbwSpace),
		theme:     gelook.NewDuoUItheme(),
		topMenuList: &layout.List{
			Axis: layout.Horizontal,
		},
		mainList: &layout.List{
			Axis: layout.Horizontal,
		},
		detailsList: &layout.List{
			Axis: layout.Horizontal,
		},
		cursor: &DuoUIfbwThing{},
	}
}

func main() {
	gofont.Register()
	fbw := NewDuoUIfbw()
	fbw.allSpaces.NewDuoUIspace(0, "/")
	//root := newDuoUIfbwSpace("/")
	//root.getThingsFromSpace()
	//col := color.RGBA{A: 0xff, R: 0x30, G: 0xcf, B: 0xcf}
	//fbw.spaces = append(fbw.spaces, root)
	fbw.headerPath = fbw.cursor.Name
	//itemValue := item{
	//	i: 0,
	//}
	//list := &layout.List{
	//	Axis: layout.Horizontal,
	//}
	go func() {
		//NewWindow(options ...Option) *Window {
		w := app.NewWindow(
			app.Size(unit.Dp(900), unit.Dp(556)),
			app.Title("Learn Gio"),
		)

		//th := material.NewTheme()

		file := &Button{
			Name: "File",
			Do: func(interface{}) {
				//itemValue.doReset()
			},
			OperateValue:   0,
			ColorBg:        "ff303030",
			ColorBgHover:   "ffcf3030",
			ColorText:      "ffcfcfcf",
			ColorTextHover: "ffcfcfcf",
		}
		edit := &Button{
			Name: "Edit",
			Do: func(interface{}) {
				//itemValue.doReset()
			},
			OperateValue:   0,
			ColorBg:        "ff303030",
			ColorBgHover:   "ffcf3030",
			ColorText:      "ffcfcfcf",
			ColorTextHover: "ffcfcfcf",
		}

		help := &Button{
			Name: "Help",
			Do: func(interface{}) {
				//itemValue.doReset()
			},
			OperateValue:   0,
			ColorBg:        "ff303030",
			ColorBgHover:   "ffcf3030",
			ColorText:      "ffcfcfcf",
			ColorTextHover: "ffcfcfcf",
		}

		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)
				layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func() {
						cs := gtx.Constraints
						DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, HexARGB("ff303030"), [4]float32{0, 0, 0, 0}, unit.Dp(0))

						in := layout.UniformInset(unit.Dp(0))
						in.Layout(gtx, func() {
							//th.H6(fmt.Sprint(itemValue.i)).Layout(gtx)
							widgets := []func(){
								func() {
									layout.Flex{}.Layout(gtx,
										layout.Rigid(func() {
											file.Layout(gtx, fbw.theme)
										}),
										layout.Rigid(func() {
											edit.Layout(gtx, fbw.theme)
										}),
										layout.Rigid(func() {
											help.Layout(gtx, fbw.theme)
										}),
									)

								},

								//func() {
								//	th.H6(fmt.Sprint(itemValue.i)).Layout(gtx)
								//},
							}
							fbw.topMenuList.Layout(gtx, len(widgets), func(i int) {
								layout.UniformInset(unit.Dp(0)).Layout(gtx, widgets[i])
							})

						})
					}),
					layout.Flexed(1, func() {
						cs := gtx.Constraints
						DrawRectangle(
							gtx,
							cs.Width.Max,
							cs.Height.Max,
							HexARGB("ffcfcfcf"),
							[4]float32{0, 0, 0, 0},
							unit.Dp(0),
						)

						layout.Flex{
							Axis: layout.Vertical,
						}.Layout(gtx,
							layout.Rigid(func() {
								fbw.theme.DuoUIlabel(
									unit.Dp(24),
									fbw.headerPath,
								).Layout(gtx)
							}),
							layout.Flexed(1, func() {
								layout.Flex{
									Axis: layout.Horizontal,
								}.Layout(gtx,
									layout.Flexed(1, func() {
										fbw.mainList.Layout(
											gtx,
											len(fbw.spaces()),
											func(i int) {
												fmt.Println("indjrx:", i)
												layout.UniformInset(unit.Dp(0)).Layout(
													gtx,
													fbw.spaces()[i].listThings(
														&fbw.cursor,
														&fbw.allSpaces,
														gtx,
														fbw.theme))
											})
									}),
									layout.Rigid(func() {
										DrawRectangle(gtx, cs.Width.Max, cs.Height.Max, HexARGB("ffcf8030"), [4]float32{0, 0, 0, 0}, unit.Dp(0))

										details := []func(){
											func() {
												fbw.theme.DuoUIlabel(unit.Dp(14), fbw.cursor.Name).Layout(gtx)
											},
											func() {
												fbw.theme.DuoUIlabel(unit.Dp(14), fbw.cursor.Type).Layout(gtx)
											},
											//func() {
											//	fbw.theme.Label(unit.Dp(14), fbw.cursor.details.filename).Layout(gtx)
											//},
											//func() {
											//	fbw.theme.Label(unit.Dp(14), fbw.cursor.details.fullPath).Layout(gtx)
											//},
											//func() {
											//	fbw.theme.Label(unit.Dp(14), fbw.cursor.details.ext).Layout(gtx)
											//},
											//func() {
											//	fbw.theme.Label(unit.Dp(14), fmt.Sprint(fbw.cursor.details.Size)).Layout(gtx)
											//},
											func() {
												//fbw.theme.Label(unit.Dp(14), fmt.Sprint(fbw.cursor.details.ModTime)).Layout(gtx)
											},
										}
										fbw.detailsList.Layout(gtx, len(details), func(i int) {
											layout.UniformInset(unit.Dp(0)).Layout(gtx, details[i])
										})
									}))
							}))

					}),
				)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
