package domain

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
	ServerTime int64
}

type Events []Event

func (extEvent *Event) Enrichment(ip string, serverTime int64) *Event {
	extEvent.Ip = ip
	extEvent.ServerTime = serverTime
	return extEvent
}