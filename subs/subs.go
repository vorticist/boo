package subs

import "gitlab.com/vorticist/logger"

type EventHandler func(e Event) error

type Event struct {
	Type string
	Data map[string]interface{}
}

var EventChannel chan Event
var subscriptions map[string]EventHandler

func NewEventListener() chan Event {
	EventChannel = make(chan Event)
	subscriptions = make(map[string]EventHandler)
	go listen()
	return EventChannel
}

func Subscribe(eventType string, handler EventHandler) {
	if subscriptions == nil {
		subscriptions = make(map[string]EventHandler)
	}
	subscriptions[eventType] = handler
}

func listen() {
	defer close(EventChannel)
	for {
		select {
		case event := <-EventChannel:
			if handler, ok := subscriptions[event.Type]; ok {
				go func(h EventHandler) {
					err := h(event)
					if err != nil {
						logger.Errorf("error handling event: %s", err)
					}
				}(handler)
			}
		}
	}
}

const (
	GetEntries      = "get_entries"
	EntriesReceived = "entries_received"
	SaveNewEntry    = "save_new_entry"
	RemoveEntry     = "remove_entry"
	FilterEntries   = "filter_entries"
)
