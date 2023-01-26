package ui

import (
	"image/color"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gitlab.com/vorticist/book/ui/assets"
	"gitlab.com/vorticist/logger"
)

func StartApp() {
	assets.Load()
	go func() {
		w := app.NewWindow(
			app.Title("Book of Omens"),
			app.Size(unit.Dp(350), unit.Dp(350)),
		)

		if err := draw(w); err != nil {
			logger.Errorf("draw error: %v", err)
		}
	}()

	app.Main()
}

func draw(w *app.Window) error {
	var ops op.Ops
	th := material.NewTheme(assets.Fonts)
	th.Palette = material.Palette{
		Fg:         color.NRGBA{R: 0xd5, G: 0x00, B: 0x37, A: 0xFF},
		Bg:         color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
		ContrastBg: color.NRGBA{R: 0xd5, G: 0x00, B: 0x37, A: 0xFF},
		ContrastFg: color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
	}
	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				mainScreen(gtx, th)
				e.Frame(gtx.Ops)
			case system.DestroyEvent:
				logger.Infof("destroy: %v", e)
				return e.Err

			}

		}

	}
}

func layoutLabel(th *material.Theme, text string) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(unit.Dp(3)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return material.Label(th, unit.Sp(16), text).Layout(gtx)
		})
	}
}

func layoutTextField(th *material.Theme, textEditor *widget.Editor, hint string) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		e := material.Editor(th, textEditor, hint)
        cornerRadius := 5
		border := widget.Border{Color: color.NRGBA{A: 0xff}, CornerRadius: unit.Dp(cornerRadius), Width: unit.Dp(1)}
		return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(unit.Dp(8)).Layout(gtx, e.Layout)
		})
	}
}
