package item

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/fbw/pkg/lyt"
	"github.com/gioapp/gel/helper"
	"image"
)

type Item struct {
	Theme      *material.Theme
	Button     *widget.Clickable
	Check      *widget.Bool
	Name       string
	Type       string
	Size       int64
	IconFolder *widget.Icon
	IconDots   *widget.Icon
}

func ItemBtn(t *material.Theme, b *widget.Clickable, c *widget.Bool, iconFolder, iconDots *widget.Icon, name, fileType string, size int64) Item {
	return Item{
		Theme:      t,
		Button:     b,
		Check:      c,
		Name:       name,
		Type:       fileType,
		Size:       size,
		IconFolder: iconFolder,
		IconDots:   iconDots,
	}
}

func (b Item) Layout(gtx layout.Context) layout.Dimensions {
	return lyt.Format(gtx, "hflexb(middle,r(_),f(0.8,_),r(_),f(0.2,_),r(_))",
		func(gtx layout.Context) layout.Dimensions {
			return material.CheckBox(b.Theme, b.Check, "").Layout(gtx)
		},
		func(gtx layout.Context) layout.Dimensions {
			bb := material.ButtonLayout(b.Theme, b.Button)
			bb.CornerRadius = unit.Dp(0)
			bb.Background = helper.HexARGB("ffffffff")
			return bb.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(4)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					gtx.Constraints.Min.X = gtx.Constraints.Max.X
					iconAndLabel := layout.Flex{Alignment: layout.Middle, Spacing: layout.SpaceBetween}
					layIcon := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							//size := gtx.Px(b.IconSize) - 2*gtx.Px(unit.Dp(16))
							b.IconFolder.Color = b.Theme.Color.Text
							b.IconFolder.Layout(gtx, unit.Px(float32(32)))
							return layout.Dimensions{
								Size: image.Point{X: 32, Y: 32},
							}
						})
					})
					layLabel := layout.Rigid(func(gtx layout.Context) layout.Dimensions {

						return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
							func(gtx layout.Context) layout.Dimensions {
								name := material.Body1(b.Theme, b.Name)
								name.Alignment = text.Start
								return name.Layout(gtx)
							},
							func(gtx layout.Context) layout.Dimensions {
								fileType := material.Caption(b.Theme, fmt.Sprint("Type: ", b.Type))
								fileType.Alignment = text.Start
								return fileType.Layout(gtx)
							},
						)
					})
					return iconAndLabel.Layout(gtx, layIcon, layLabel)
				})
			})
		},
		func(gtx layout.Context) layout.Dimensions {
			return material.Body1(b.Theme, "0").Layout(gtx)
		},
		func(gtx layout.Context) layout.Dimensions {
			return material.Body1(b.Theme, fmt.Sprint(b.Size)).Layout(gtx)
		},
		func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				//size := gtx.Px(b.IconSize) - 2*gtx.Px(unit.Dp(16))
				b.IconDots.Color = b.Theme.Color.Text
				b.IconDots.Layout(gtx, unit.Px(float32(32)))
				return layout.Dimensions{
					Size: image.Point{X: 32, Y: 32},
				}
			})
		},
	)
}
