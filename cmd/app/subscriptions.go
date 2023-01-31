package main

import (
	"github.com/vorticist/boo/client"
	"github.com/vorticist/boo/subs"
)

func subscriptions(cli client.Client) {
	subs.Subscribe(subs.GetEntries, func(e subs.Event) error {
		entries, err := cli.GetEntries()
		if err != nil {
			return err
		}

		var pairs [][]string
		for key, value := range entries {
			pairs = append(pairs, []string{key, value})
		}

		subs.EventChannel <- subs.Event{
			Type: subs.EntriesReceived,
			Data: map[string]interface{}{
				"entries": pairs,
			},
		}

		return nil
	})
}
