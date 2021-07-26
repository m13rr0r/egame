package domain

import "time"

type Event struct {
	ClientTime string `json:"client_time"`
	DeviceId   string `json:"device_id"`
	DeviceOs   string `json:"device_os"`
	Session    string `json:"session"`
	Sequence   int    `json:"sequence"`
	Event      string `json:"event"`
	ParamInt   int    `json:"param_int"`
	ParamStr   string `json:"param_str"`
	Ip         string
	ServerTime time.Time
}

func (extEvent *Event) Enrichment(ip string, serverTime time.Time) *Event {
	extEvent.Ip = ip
	extEvent.ServerTime = serverTime
	return extEvent
}