package handler

import (
	"time"

	"github.com/m13rr0r/egame/infrastructure/persistence"
	"github.com/m13rr0r/egame/infrastructure/storage"
	"github.com/valyala/fastjson"
)

func (c *Context) TimeLimitHandler(eventChan chan *fastjson.Value, stg *storage.Storage) {
	for {
		if len(eventChan) > 0 {
			lenChan := len(eventChan)
			go persistence.PutEvents(stg, eventChan, lenChan)
		}
		time.Sleep(time.Second * 1)
	}
}

func (c *Context) MaxLimitHandler(eventChan chan *fastjson.Value, stg *storage.Storage) {
	for {
		if len(eventChan) == c.Limit {
			go persistence.PutEvents(stg, eventChan, c.Limit)
		}
	}
}
