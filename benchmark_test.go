package server

import (
	"fmt"
	"sync"
	"testing"
)

func testOne(m *manager) error {
	x := "hello"
	for i := 0; i < 10; i++ {
		x = x + x
	}
	_ = x
	r := Request{}
	rep := m.Handle(r)
	if want, have := 5<<10, len(rep.A); have != want {
		return fmt.Errorf("want %d, have %d", want, have)
	}
	return nil
}

func TestOne(t *testing.T) {
	testOne(NewManager())
}

func runBench(b *testing.B, n int) {
	m := NewManager()

	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		for j := 0; j < n; j++ {
			wg.Add(1)
			go func() {
				if err := testOne(m); err != nil {
					b.Error(err)
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func Benchmark100(b *testing.B) {
	runBench(b, 100)
}
