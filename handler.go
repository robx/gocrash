package server

type hrequest struct {
	req Request
	rep chan<- Response
}

type handler struct {
	requests chan hrequest
}

func NewHandler() *handler {
	h := handler{
		requests: make(chan hrequest),
	}
	go h.loop()
	return &h
}

func (h *handler) loop() {
	for r := range h.requests {
		r.rep <- run(r.req)
	}
}

func (h *handler) Handle(req Request, rep chan<- Response) {
	r := hrequest{req, rep}
	h.requests <- r
}

func run(req Request) Response {
	x := "world"
	for i := 0; i < 10; i++ {
		x = x + x
	}
	return Response{A: x}
}
