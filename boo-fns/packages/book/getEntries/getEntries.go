package main

import (
	"encoding/json"
	"github.com/vorticist/boo/storage"
	"gitlab.com/vorticist/logger"
)

func Main(args map[string]interface{}) map[string]interface{} {
	stor, err := storage.NewStorer()
	if err != nil {
		return nil
	}
	entries := stor.GetEntries()
	entriesJson, err := json.Marshal(entries)
	if err != nil {
		logger.Errorf("failed to marshal entries: %v", err)
		return nil
	}
	jsonString := string(entriesJson)
	return map[string]interface{}{
		"entries": jsonString,
	}
}
