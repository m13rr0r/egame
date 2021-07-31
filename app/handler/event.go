package handler

import (
	"bufio"
	"net/http"

	"github.com/valyala/fastjson"
)

func (c *Context) CreateEvent(w http.ResponseWriter, r *http.Request) {

	scanner := bufio.NewScanner(r.Body)

	for scanner.Scan() {
		var p fastjson.Parser
		value, err := p.Parse(scanner.Text())
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		} else {
			c.EventChan <- value
		}
	}
	 w.WriteHeader(http.StatusOK)
}

