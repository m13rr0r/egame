package handler

import (
	"bufio"
	"bytes"
	"fmt"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/m13rr0r/egame/domain"
	"github.com/valyala/fasthttp"
)

func (c *Context)CreateEvent(ctx *fasthttp.RequestCtx) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	scanner := bufio.NewScanner(bytes.NewReader(ctx.Request.Body()))
	for scanner.Scan() {
		event := c.EventPool.Get().(*domain.Event)
		err := json.Unmarshal(scanner.Bytes(), &event)
		if err != nil {
			fmt.Fprintf(ctx, "%v\n", err)
			ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}
		c.EventChan <- event.Enrichment(getIp(), getServerTime())
	}

	ctx.Response.SetStatusCode(fasthttp.StatusOK)
}


func getIp() string {
	return "8.8.8.8"
}

func getServerTime() int64 {
	return time.Now().Unix()
}

