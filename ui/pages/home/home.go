package home

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/vorticist/boo/subs"
	"github.com/vorticist/boo/ui/pages"
	"gitlab.com/vorticist/logger"
)

var (
	list *layout.List = &layout.List{
		Axis: layout.Vertical,
	}
	entries [][]string
)

func Page(th *material.Theme) pages.Page {
	return &homePage{
		Theme: th,
	}
}

type homePage struct {
	layout.List
	Theme *material.Theme
}

func (m *homePage) Start() {
	logger.Info("Home page started")
	subs.EventChannel <- subs.Event{
		Type: subs.GetEntries,
		Data: nil,
	}

	subs.Subscribe(subs.EntriesReceived, func(e subs.Event) error {
		entries = e.Data["entries"].([][]string)
		return nil
	})
}

func (m *homePage) Layout(gtx layout.Context) layout.Dimensions {
	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Vertical,
	}.Layout(gtx,
		layout.Flexed(1, m.layoutItems),
	)
}

func (m *homePage) Events() {
	logger.Info("home page subs")
}

func (m *homePage) layoutItems(gtx layout.Context) layout.Dimensions {
	l := list
	return l.Layout(gtx, len(entries), func(gtx layout.Context, i int) layout.Dimensions {
		in := layout.Inset{}
		switch i {
		case 0:
			in.Top = unit.Dp(4)
		case len(entries) - 1:
			in.Bottom = unit.Dp(4)
		}
		return in.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return m.itemLayout(gtx, i)
		})
	})
}
func (m *homePage) itemLayout(gtx layout.Context, i int) layout.Dimensions {
	in := layout.Inset{
		Top:    unit.Dp(16),
		Bottom: unit.Dp(16),
		Left:   unit.Dp(16),
		Right:  unit.Dp(16),
	}
	entry := entries[i]
	// click := &p.clicks[i]
	logger.Infof("entry[%v]: [%v]", i, entry)
	dims := in.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
			layout.Flexed(1, material.Caption(m.Theme, entry[0]).Layout),
		)
	})
	// click.Add(gtx.Ops)
	return dims
}
