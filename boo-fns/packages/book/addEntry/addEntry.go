package main

import "github.com/vorticist/boo/storage"

func Main(args map[string]interface{}) map[string]string {
	stor, err := storage.NewStorer()
	if err != nil {
		return nil
	}
	entries := stor.GetEntries()
	return entries
}
