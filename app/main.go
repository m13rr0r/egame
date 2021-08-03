package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/m13rr0r/egame/domain"
	"github.com/m13rr0r/egame/handler"
	"github.com/m13rr0r/egame/infrastructure/storage"
	"github.com/valyala/fasthttp"
)

func main() {
	chanSize, _ := strconv.Atoi(os.Getenv("APP_CHAN_SIZE"))
	chunkSize, _ := strconv.Atoi(os.Getenv("APP_CHUNK_SIZE"))

	stg, err := storage.NewStorage(os.Getenv("APP_DB_STORAGE_URI"))
	if err != nil {
		fmt.Printf("%v\n", err.Error())
	}

	eventPool := &sync.Pool {
		New: func()interface{} {
			return new(domain.Event)
		},
	}

	eventChan :=  make (chan *domain.Event, chanSize)

	c := &handler.Context{
		Storage: stg,
		ChanSize: chanSize,
		ChunkSize: chunkSize,
		EventPool: eventPool,
		EventChan: eventChan,
	}

	go c.TimeLimitHandler()
	go c.MaxLimitHandler()

	fasthttp.ListenAndServe(":80", c.CreateEvent)
}