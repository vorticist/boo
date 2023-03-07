package storage

import (
	"gitlab.com/vorticist/logger"
	"testing"
)

func TestNewStorer(t *testing.T) {
	storer, err := NewStorer()
	if err != nil || storer == nil {
		t.Errorf("NewStorer failed: %v", err)
	}
}

func TestStorer_GetEntries(t *testing.T) {
	storer, err := NewStorer()
	if err != nil {
		t.Errorf("TestStorer_GetEntries -> NewStorer failed: %v", err)
		return
	}
	entries := storer.GetEntries()
	if entries == nil {
		t.Errorf("TestStorer_GetEntries -> entries = nil")
		return
	}

	logger.Infof("entries: %v", entries)
}

func TestStorer_SaveEntry(t *testing.T) {
	storer, err := NewStorer()
	if err != nil {
		t.Errorf("TestStorer_GetEntries -> NewStorer failed: %v", err)
		return
	}
	entries := storer.GetEntries()

	storer.SaveEntry("NEW_KEY", "NEW_VALUE")

	entries = storer.GetEntries()

	if _, ok := entries["NEW_KEY"]; !ok {
		t.Error("TestStorer_AddEntry -> key not added")
		return
	}

	if entries["NEW_KEY"] != "NEW_VALUE" {
		t.Error("TestStorer_AddEntry -> value not stored")
	}
}

func TestStorer_RemoveEntry(t *testing.T) {
	storer, err := NewStorer()
	if err != nil {
		t.Errorf("TestStorer_GetEntries -> NewStorer failed: %v", err)
		return
	}

	storer.RemoveEntry("hotmail")

	entries := storer.GetEntries()

	if _, ok := entries["hotmail"]; ok {
		t.Error("TestStorer_AddEntry -> key not removed")
		return
	}
}
