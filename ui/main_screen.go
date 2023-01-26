package ui

import (
	"gioui.org/layout"
    "gioui.org/unit"
    "gioui.org/widget"
    "gioui.org/widget/material"
)

var (
    list = &widget.List{
        List: layout.List{
            Axis: layout.Vertical,
            },
            }
)

func mainScreen(gtx layout.Context, th *material.Theme) {
	layout.Flex{
		Axis:    layout.Vertical,
		Spacing: layout.SpaceAround,
	}.Layout(gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{
				Axis: layout.Horizontal,
			}.Layout(gtx,
				layout.Flexed(1, layoutControlList(th)),
			)
		}),
	)
}

func layoutControlList(th *material.Theme) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
//        in := layout.UniformInset(unit.Dp(8))
        widgets := []layout.Widget{

        }

         	return material.List(th, list).Layout(gtx, len(widgets), func(gtx layout.Context, index int) layout.Dimensions {
         		return layout.UniformInset(unit.Dp(16)).Layout(gtx, widgets[index])
         	})
    }
}
