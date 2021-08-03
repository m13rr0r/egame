package handler

import (
	"time"

	"github.com/m13rr0r/egame/domain"
	"github.com/m13rr0r/egame/infrastructure/persistence"
)

func (c *Context) TimeLimitHandler() {
	for {
		if len(c.EventChan) > 0 {
			var events []*domain.Event
			for i := 0; i < c.ChunkSize; i++ {
				if len(c.EventChan) == 0 {
					break
				}
				event := <-c.EventChan
				events = append(events, event)
				c.EventPool.Put(event)
			}
			if len(events) > 0 {
				go persistence.PutEvents(c.Storage, events)
			}
		}
		time.Sleep(time.Second * 1)
	}
}

func (c *Context) MaxLimitHandler() {
	for {
		if len(c.EventChan) >= c.ChunkSize {
			var events []*domain.Event
			for i := 0; true ; i++ {
				if len(c.EventChan) == 0 {
					break
				}
				event := <-c.EventChan
				events = append(events, event)
				c.EventPool.Put(event)
				if len(events) == c.ChunkSize {
					go persistence.PutEvents(c.Storage, events)
					events = []*domain.Event{}
					break
				}
			}
			if len(events) > 0 {
				go persistence.PutEvents(c.Storage, events)
			}
		}
	}
}

