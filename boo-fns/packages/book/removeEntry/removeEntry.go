package main

import (
	"fmt"
	"github.com/vorticist/boo/storage"
)

func Main(args map[string]interface{}) map[string]interface{} {
	var keyToRemove string
	if _, ok := args["key"]; !ok {
		return map[string]interface{}{
			"error": "key not present",
		}
	}
	keyToRemove = fmt.Sprintf("%v", args["key"])

	stor, err := storage.NewStorer()
	if err != nil {
		return nil
	}
	stor.RemoveEntry(keyToRemove)
	entries := stor.GetEntries()
	return map[string]interface{}{
		"entries": entries,
	}
}
