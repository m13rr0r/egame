package persistence

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/m13rr0r/egame/domain"
)

func PutEvent(stmt *sql.Stmt, logger echo.Logger, event *domain.Event) error {
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
	return nil
}

func GetPutEventPrepare(tx *sql.Tx) (*sql.Stmt, error) {
	stmt, err := tx.Prepare("INSERT INTO egame.events (client_time, device_id, device_os, session, sequence, event, " +
		"param_int, param_str, ip, server_time) VALUES (?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return stmt, err
	}

	return stmt, nil
}
