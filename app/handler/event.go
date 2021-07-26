package handler

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/m13rr0r/egame/domain"
	"github.com/m13rr0r/egame/infrastructure/persistence"
)

const ip = "8.8.8.8"

func (h *Handler) CreateEvent(c echo.Context) error {
	tx, err := h.DB.Client.Begin()
	if err != nil {
		c.String(http.StatusBadRequest, "")
	}
	for err := range process(tx, c.Logger(), c.Request().Body) {
		if err != nil {
			c.Logger().Printf("%v\n", err)
			return c.String(http.StatusBadRequest, "")
		}
	}
	err = tx.Commit()
	if err != nil {
		return c.String(http.StatusBadGateway, "")
	}
	return c.String(http.StatusOK, "")
}

func process(tx *sql.Tx, logger echo.Logger, body io.ReadCloser) chan error {
	ch := make(chan error)
	go func() {
		var wg sync.WaitGroup
		scanner := bufio.NewScanner(body)
		stmt, _ := persistence.GetPutEventPrepare(tx)
		defer stmt.Close()
		for scanner.Scan() {
			wg.Add(1)
			text := scanner.Text()
			go func(stmt *sql.Stmt, logger echo.Logger, text string) {
				defer wg.Done()
				err := save(stmt, logger, text)
				if err != nil {
					ch <- err
				}

			}(stmt, logger, text)
		}
		wg.Wait()
		close(ch)
	}()

	return ch
}

func save(stmt *sql.Stmt, logger echo.Logger, eventString string) error {
	event := new(domain.Event)
	err := match(eventString, event)
	if err != nil {
		return err
	}

	event.Enrichment(getIp(),getServerTime())
	err = persistence.PutEvent(stmt, logger, event)
	if err != nil {
		return err
	}

	return nil
}

func match(eventString string, event *domain.Event) error {
	return json.Unmarshal([]byte(eventString), event)
}

func getIp() string {
	return ip
}

func getServerTime() time.Time {
	return time.Now()
}
