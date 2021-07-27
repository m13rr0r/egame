package persistence

import (
	"github.com/labstack/echo/v4"
	"github.com/m13rr0r/egame/domain"
	"github.com/m13rr0r/egame/infrastructure/storage"
)

func PutEvents(db *storage.Storage, logger echo.Logger, events *domain.Events) error {
	tx, _ := db.Client.Begin()
	stmt, _ := tx.Prepare("INSERT INTO egame.events (client_time, device_id, device_os, session, sequence, event, " +
		"param_int, param_str, ip, server_time) FORMAT VALUES (?,?,?,?,?,?,?,?,?,?)")
	defer stmt.Close()

	for _, event := range *events {
		if _, err := stmt.Exec(
			event.ClientTime,
			event.DeviceId,
			event.DeviceOs,
			event.Session,
			event.Sequence,
			event.Event,
			event.ParamInt,
			event.ParamStr,
			event.Ip,
			event.ServerTime,
		); err != nil {
			logger.Printf("%v\n", err)
			return err
		}
	}
	err := tx.Commit()
	if err != nil {
		logger.Printf("%v\n", err)
		return err
	}

	return nil
}
