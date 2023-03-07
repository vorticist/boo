package main

import (
	"gioui.org/widget"
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

		es := mapEntries(entries)

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
		entry.Editing = false
		entries, err := cli.SaveEntry(entry.Key, entry.Value)
		if err != nil {
			logger.Errorf("error saving entry: %v", err)
			return err
		}

		es := mapEntries(entries)

		subs.EventChannel <- subs.Event{
			Type: subs.EntriesReceived,
			Data: map[string]interface{}{
				"entries": es,
			},
		}

		return nil
	})

	subs.Subscribe(subs.RemoveEntry, func(e subs.Event) error {
		entry := e.Data["entry"].(models.Entry)
		entries, err := cli.RemoveEntry(entry.Key)
		if err != nil {
			logger.Errorf("error removing entry: %v", err)
			return err
		}

		es := mapEntries(entries)

		subs.EventChannel <- subs.Event{
			Type: subs.EntriesReceived,
			Data: map[string]interface{}{
				"entries": es,
			},
		}

		return nil
	})
}

func mapEntries(e map[string]string) []models.Entry {
	var es []models.Entry
	for key, value := range e {
		es = append(es, models.Entry{
			Key:       key,
			Value:     value,
			Editing:   false,
			ShowBtn:   new(widget.Clickable),
			DeleteBtn: new(widget.Clickable),
			CopyBtn:   new(widget.Clickable),
		})
	}

	return es
}
