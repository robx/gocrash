package server

type hrequest struct {
	done func(Response)
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
		r.done(run())
	}
}

func (h *handler) Handle(done func(Response)) {
	r := hrequest{done}
	h.requests <- r
}

func run() Response {
	x := "world"
	for i := 0; i < 10; i++ {
		x = x + x
	}
	return Response{A: x}
}
