package handler

import (
	"bufio"
	"encoding/json"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/m13rr0r/egame/domain"
	"github.com/m13rr0r/egame/infrastructure/persistence"
	"github.com/m13rr0r/egame/infrastructure/storage"
)

const ip = "8.8.8.8"

func (h *Handler) CreateEvent(c echo.Context) error {
	events := new(domain.Events)
	for err := range process(c.Logger(), c.Request().Body, events) {
		if err != nil {
			c.Logger().Printf("%v\n", err)
			return c.String(http.StatusBadRequest, "")
		}
	}

	err := save(h.DB, c.Logger(), events)
	if err != nil {
		return c.String(http.StatusBadRequest, "")
	}

	return c.String(http.StatusOK, "")
}

func process(logger echo.Logger, body io.ReadCloser, events *domain.Events) chan error {
	ch := make(chan error)
	go func() {
		var wg sync.WaitGroup
		scanner := bufio.NewScanner(body)
		for scanner.Scan() {
			wg.Add(1)
			text := scanner.Text()
			go func(logger echo.Logger, events *domain.Events, text string) {
				defer wg.Done()
				event, err := match(text)
				if err != nil {
					ch <- err
				} else {
					event.Enrichment(getIp(),getServerTime())
					*events = append(*events, *event)

				}
			}(logger, events, text)
		}
		wg.Wait()
		close(ch)
	}()

	return ch
}

func save(db *storage.Storage, logger echo.Logger, events *domain.Events) error {
	err := persistence.PutEvents(db, logger, events)
	if err != nil {
		return err
	}

	return nil
}

func match(eventString string) (*domain.Event, error) {
	event := new(domain.Event)
	err := json.Unmarshal([]byte(eventString), event)
	return event, err
}

func getIp() string {
	return ip
}

func getServerTime() int64 {
	return time.Now().Unix()
}
