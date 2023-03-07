package models

import "gioui.org/widget"

type Entry struct {
	Key          string
	Value        string
	Editing      bool
	ShowPassword bool
	ShowBtn      *widget.Clickable
	DeleteBtn    *widget.Clickable
	CopyBtn      *widget.Clickable
}
