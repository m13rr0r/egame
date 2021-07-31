package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/m13rr0r/egame/handler"
	"github.com/m13rr0r/egame/infrastructure/storage"
	"github.com/valyala/fastjson"
)

// LimitSize https://clickhouse.tech/docs/ru/introduction/performance/
const LimitSize = 500000

func main() {

	stg, err := storage.NewStorage(os.Getenv("APP_DB_STORAGE_URI"))
	if err != nil {
		fmt.Printf("%v\n", err.Error())
	}

	eventChan := make(chan *fastjson.Value, LimitSize)

	c := &handler.Context{Limit: LimitSize, EventChan: eventChan}

	// Workers
	go c.TimeLimitHandler(eventChan, stg)

	go c.MaxLimitHandler(eventChan, stg)

	// Routes
	http.HandleFunc("/event", c.CreateEvent)

	// Start server
	err = http.ListenAndServe(":80", nil)
	if err != nil {
		return
	} else {
		fmt.Println("Started at :8081")
	}
}
