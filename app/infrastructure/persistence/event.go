package persistence

import (
	"fmt"
	"time"

	"github.com/m13rr0r/egame/infrastructure/storage"
	"github.com/valyala/fastjson"
)

func PutEvents(db *storage.Storage, events chan *fastjson.Value, limit int) {
	fmt.Printf("%v\n", "+1")
	tx, _ := db.Client.Begin()
	stmt, _ := tx.Prepare("INSERT INTO egame.events_buffer (client_time, device_id, device_os, session, sequence, event, " +
		"param_int, param_str, ip, server_time) FORMAT VALUES (?,?,?,?,?,?,?,?,?,?)")
	defer stmt.Close()

	for i := 0; i <= limit; i++ {
		event := <- events
		if _, err := stmt.Exec(
			event.GetStringBytes("client_time"),
			event.GetStringBytes("device_id"),
			event.GetStringBytes("device_os"),
			event.GetStringBytes("session"),
			event.GetInt64("sequence"),
			event.GetStringBytes("event"),
			event.GetInt64("param_int"),
			event.GetStringBytes("param_str"),
			getIp(),
			getServerTime(),
		); err != nil {
			fmt.Printf("%v\n", err.Error())
		}
	}
	err := tx.Commit()
	if err != nil {
		fmt.Printf("%v\n", err.Error())
	}
}

func getIp() string {
	return "8.8.8.8"
}

func getServerTime() int64 {
	return time.Now().Unix()
}
