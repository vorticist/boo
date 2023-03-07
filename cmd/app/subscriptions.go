package main

import (
	"github.com/vorticist/boo/client"
	"github.com/vorticist/boo/models"
	"github.com/vorticist/boo/subs"
	"gitlab.com/vorticist/logger"
)

func subscriptions(cli client.Client) {
	subs.Subscribe(subs.GetEntries, func(e subs.Event) error {
		entries, err := cli.GetEntries()
		if err != nil {
			return err
		}

		var es []models.Entry
		for key, value := range entries {
			es = append(es, models.Entry{
				Key:     key,
				Value:   value,
				Editing: false,
			})
		}

		subs.EventChannel <- subs.Event{
			Type: subs.EntriesReceived,
			Data: map[string]interface{}{
				"entries": es,
			},
		}

		return nil
	})

	subs.Subscribe(subs.SaveNewEntry, func(e subs.Event) error {
		entry := e.Data["entry"].(models.Entry)
		entries, err := cli.SaveEntry(entry.Key, entry.Value)
		if err != nil {
			logger.Errorf("error saving entry: %v", err)
			return err
		}

		subs.EventChannel <- subs.Event{
			Type: subs.EntriesReceived,
			Data: map[string]interface{}{
				"entries": entries,
			},
		}

		return nil
	})
}
