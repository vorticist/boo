package client

import (
	"gitlab.com/vorticist/logger"
	"testing"
)

func TestClient_GetEntries(t *testing.T) {
	c := New()
	entries, err := c.GetEntries()
	if err != nil {
		t.Errorf("failed to get entries -> %v", err)
	}

	logger.Infof("entries: %v", entries)
}

func TestClient_SaveEntry(t *testing.T) {
	c := New()
	entries, err := c.SaveEntry("API_KEY_THIS", "api this value")
	if err != nil {
		t.Errorf("failed to get entries -> %v", err)
		return
	}

	if _, ok := entries["API_KEY_THIS"]; !ok {
		t.Errorf("key not found")
		return
	}

	if "api this value" != entries["API_KEY_THIS"] {
		t.Errorf("expected API_KEY_THIS value is not correct")
		return
	}

	logger.Infof("entries: %v", entries)
}

func TestClient_RemoveEntry(t *testing.T) {
	c := New()
	entries, err := c.RemoveEntry("API_KEY_THIS")
	if err != nil {
		t.Errorf("failed to remove entry -> %v", err)
	}

	if _, ok := entries["API_KEY_THIS"]; ok {
		t.Errorf("key not removed")
		return
	}

	logger.Infof("entries: %v", entries)
}
