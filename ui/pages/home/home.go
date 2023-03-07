package home

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/vorticist/boo/crypt"
	"github.com/vorticist/boo/models"
	"github.com/vorticist/boo/subs"
	"github.com/vorticist/boo/ui/pages"
	"gitlab.com/vorticist/logger"
	"golang.org/x/exp/shiny/materialdesign/icons"
)

var (
	list *layout.List = &layout.List{
		Axis: layout.Vertical,
	}
	entries []models.Entry
)

func Page(th *material.Theme, env *pages.Env) pages.Page {
	ic, err := widget.NewIcon(icons.ContentAdd)
	if err != nil {
		logger.Error(err)
	}
	icc, err := widget.NewIcon(icons.NavigationCancel)
	if err != nil {
		logger.Error(err)
	}
	return &homePage{
		Theme:     th,
		env:       env,
		addBtn:    new(widget.Clickable),
		saveBtn:   new(widget.Clickable),
		cancelBtn: new(widget.Clickable),
		passEditor: &widget.Editor{
			SingleLine: true,
		},
		keyEditor: &widget.Editor{
			SingleLine: true,
		},
		saveIcon:   ic,
		cancelIcon: icc,
	}
}

type homePage struct {
	layout.List
	Theme *material.Theme
	env   *pages.Env

	addBtn     *widget.Clickable
	saveBtn    *widget.Clickable
	cancelBtn  *widget.Clickable
	passEditor *widget.Editor
	keyEditor  *widget.Editor
	saveIcon   *widget.Icon
	cancelIcon *widget.Icon
}

func (m *homePage) Start() {
	logger.Info("Home page started")
	subs.EventChannel <- subs.Event{
		Type: subs.GetEntries,
		Data: nil,
	}

	subs.Subscribe(subs.EntriesReceived, func(e subs.Event) error {
		entries = e.Data["entries"].([]models.Entry)
		go func() {
			m.env.Redraw()
		}()
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
	count := len(entries) + 1
	return l.Layout(gtx, count, func(gtx layout.Context, i int) layout.Dimensions {
		in := layout.Inset{}
		switch i {
		case 0:
			in.Top = unit.Dp(4)
		case count - 1:
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

	if i == 0 {
		return in.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(layout.Spacer{Height: unit.Dp(16)}.Layout),
				layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
					for m.addBtn.Clicked() {
						logger.Info("adding")
						entries = append(entries, models.Entry{
							Key:     "New Key",
							Value:   "New Password",
							Editing: true,
						})
					}
					return material.Clickable(gtx, m.addBtn, func(gtx layout.Context) layout.Dimensions {
						flatBtnText := material.Body1(m.Theme, "Add")
						if gtx.Queue == nil {
							flatBtnText.Color.A = 150
						}
						return layout.Center.Layout(gtx, flatBtnText.Layout)
					})
				}),
			)
		})
	}
	i -= 1
	c := crypt.Get()
	entry := entries[i]
	// click := &p.clicks[i]
	logger.Infof("entry[%v]: [%v]", i, entry)
	dims := in.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		if entry.Editing {
			for m.saveBtn.Clicked() {
				encodedPass, err := c.Encrypt(m.passEditor.Text())
				if err != nil {
					logger.Errorf("failed to encode pass: %v", err)
					continue
				}
				logger.Infof("encoded pass: %v", encodedPass)
				subs.EventChannel <- subs.Event{
					Type: subs.SaveNewEntry,
					Data: map[string]interface{}{
						"entry": models.Entry{Key: m.keyEditor.Text(), Value: encodedPass},
					},
				}
			}
			return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
				layout.Flexed(2, material.Editor(m.Theme, m.keyEditor, entry.Key).Layout),
				layout.Flexed(2, material.Editor(m.Theme, m.passEditor, entry.Value).Layout),
				layout.Flexed(1, material.IconButton(m.Theme, m.saveBtn, m.saveIcon, "Save").Layout),
				layout.Flexed(1, material.IconButton(m.Theme, m.cancelBtn, m.cancelIcon, "Cancel").Layout),
			)
		}

		decoded, err := c.Decrypt(entry.Value)
		if err != nil {
			logger.Errorf("failed to decrypt password: %v", err)
			decoded = ""
		}
		return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
			layout.Flexed(1, material.Caption(m.Theme, entry.Key).Layout),
			layout.Flexed(1, material.Caption(m.Theme, decoded).Layout),
		)
	})
	// click.Add(gtx.Ops)
	return dims
}
