package persistence

import (
	"fmt"
	"time"

	"github.com/m13rr0r/egame/domain"
	"github.com/m13rr0r/egame/infrastructure/storage"
)

func PutEvents(db *storage.Storage, events  []*domain.Event) {
	tx, err := db.Client.Begin()
	if err != nil {
		fmt.Printf("%v%v%v\n", time.Now().String(), " Error ", err.Error())
		return
	}
	stmt, err := tx.Prepare("INSERT INTO egame.events_buffer (client_time, device_id, device_os, session, sequence, event, " +
		"param_int, param_str, ip, server_time) FORMAT VALUES (?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		fmt.Printf("%v%v%v\n", time.Now().String(), " Error ", err.Error())
		return
	}
	defer stmt.Close()

	for _, event := range events {
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
			fmt.Printf("%v\n", err.Error())
		}
	}
	err = tx.Commit()
	if err != nil {
		fmt.Printf("%v%v%v\n", time.Now().String(), " Error ", err.Error())
	}
}
