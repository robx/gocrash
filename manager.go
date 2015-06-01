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

type manager chan chan Response

func NewManager() manager {
	h := NewHandler()
	m := make(manager)
	go m.loop(h)
	return m
}

func (m manager) loop(h handler) {
	for r := range m {
		r2 := r
		h <- func(rs Response) {
			r2 <- rs
		}
	}
}

func (m manager) Handle() Response {
	rep := make(chan Response)
	m <- rep
	return <-rep
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
