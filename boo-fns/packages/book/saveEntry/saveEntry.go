package main

import (
	"fmt"
	"github.com/vorticist/boo/storage"
)

func Main(args map[string]interface{}) map[string]interface{} {
	var newKey, newValue string
	if _, ok := args["key"]; !ok {
		return map[string]interface{}{
			"error": "key not present",
		}
	}
	newKey = fmt.Sprintf("%v", args["key"])
	if _, ok := args["value"]; !ok {
		return map[string]interface{}{
			"error": "value not present",
		}
	}
	newValue = fmt.Sprintf("%v", args["value"])

	stor, err := storage.NewStorer()
	if err != nil {
		return nil
	}
	stor.SaveEntry(newKey, newValue)
	entries := stor.GetEntries()
	return map[string]interface{}{
		"entries": entries,
	}
}
