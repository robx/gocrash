package server

import (
	"sync"
)

type hrequest struct {
	req  Request
	done func(Response)
}

type handler struct {
	requests chan hrequest
	wg       sync.WaitGroup
}

func NewHandler() *handler {
	h := handler{
		requests: make(chan hrequest),
	}
	h.wg.Add(1)
	go func() {
		h.loop()
		h.wg.Done()
	}()
	return &h
}

func (h *handler) loop() {
	for {
		select {
		case r, ok := <-h.requests:
			if !ok {
				return
			}
			r.done(run(r.req))
		}
	}
}

func (h *handler) Handle(req Request, done func(Response)) {
	r := hrequest{req, done}
	h.requests <- r
}

func run(req Request) Response {
	x := "world"
	for i := 0; i < 10; i++ {
		x = x + x
	}
	return Response{A: x}
}
