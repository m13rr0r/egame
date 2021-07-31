package handler

import (
	"github.com/valyala/fastjson"
)

type (
	Context struct {
		Limit     int
		EventChan chan *fastjson.Value
	}
)
