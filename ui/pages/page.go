package pages

import (
	"gioui.org/layout"
)

type Page interface {
	Start()
	Layout(layout.Context) layout.Dimensions
	Events()
}

type Env struct {
	Redraw func()
	Insets layout.Inset
}
