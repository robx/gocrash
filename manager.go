package server

import (
	"sync"
)

type request struct {
	req Request
	rep chan chan Response
}

type manager struct {
	handler  *handler
	requests chan request
}

func NewManager() *manager {
	m := manager{
		requests: make(chan request),
		handler:  NewHandler(),
	}
	go m.loop()
	return &m
}

func (m *manager) loop() {
	for {
		select {
		case r := <-m.requests:
			reps := make(chan Response, 1)
			r.rep <- reps
			wg := sync.WaitGroup{}
			wg.Add(1)
			m.handler.Handle(r.req, func(rs Response) {
				reps <- rs
				wg.Done()
			})
			go func() {
				wg.Wait()
				close(reps)
			}()
		}
	}
}

func (m *manager) Handle(req Request) Response {
	rep := make(chan chan Response)
	m.requests <- request{
		req: req,
		rep: rep,
	}
	reps := <-rep
	return <-reps
}