package main

import (
	"github.com/vorticist/boo/storage"
)

func Main(args map[string]interface{}) map[string]interface{} {
	stor, err := storage.NewStorer()
	if err != nil {
		return nil
	}
	entries := stor.GetEntries()
	return map[string]interface{}{
		"entries": entries,
	}
}
