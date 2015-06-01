package server

type Response struct {
	A string
	B string
	C string
	D string
	E string
	F string
	G string
	H string
}

type manager chan chan chan Response

func NewManager() manager {
	m := make(manager)
	go m.loop(NewHandler())
	return m
}

func (m manager) loop(h handler) {
	for r := range m {
		reps := make(chan Response)
		r <- reps
		f := func(rs Response) {
			reps <- rs
		}
		h.Handle(f)
	}
}

func (m manager) Handle() Response {
	rep := make(chan chan Response)
	m <- rep
	reps := <-rep
	return <-reps
}

type handler chan func(Response)

func NewHandler() handler {
	h := make(handler)
	go h.loop()
	return h
}

func (h handler) loop() {
	for r := range h {
		x := "world"
		for i := 0; i < 10; i++ {
			x = x + x
		}
		r(Response{A: x})
	}
}

func (h handler) Handle(done func(Response)) {
	h <- done
}
