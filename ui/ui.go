package ui

import (
	pages2 "github.com/vorticist/boo/ui/pages"
	"github.com/vorticist/boo/ui/pages/home"
	"image/color"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/vorticist/boo/ui/assets"
	pgs "github.com/vorticist/boo/ui/pages"
	"gitlab.com/vorticist/logger"
)

var (
	pages     []pages2.Page
	pageIndex int = 0
	env       *pgs.Env
)

func StartApp() {
	assets.Load()
	go func() {
		w := app.NewWindow(
			app.Title("Book of Omens"),
			app.Size(unit.Dp(550), unit.Dp(350)),
		)

		if err := draw(w); err != nil {
			logger.Errorf("draw error: %v", err)
		}
	}()

	app.Main()
}

func draw(w *app.Window) error {
	var ops op.Ops
	env = &pgs.Env{Redraw: w.Invalidate}
	th := material.NewTheme(assets.Fonts)
	th.Palette = material.Palette{
		Fg:         color.NRGBA{R: 0xd5, G: 0x00, B: 0x37, A: 0xFF},
		Bg:         color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
		ContrastBg: color.NRGBA{R: 0xd5, G: 0x00, B: 0x37, A: 0xFF},
		ContrastFg: color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
	}
	pages = append(pages, home.Page(th, env))
	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {
			case system.StageEvent:
				if e.Stage >= system.StageRunning {
					page := pages[pageIndex]
					page.Start()
				}
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				page := pages[pageIndex]
				page.Layout(gtx)
				//mainScreen(gtx, th)
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
