package main

import (
	"sync"
	"testing"
)

type Queue[T any] interface {
	Len() int
	Cap() int
	SetCap(v int)
	Dequeue() (t T)
	Enqueue(t T)
}

func BufLongLivedCustomQueue[T any](in <-chan T, q Queue[T]) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		var (
			next T
			nout chan T
		)
		for in != nil || nout != nil {
			if q.Len() < q.Cap()/4 {
				q.SetCap(q.Len() * 2)
			}
			select {
			case nout <- next:
				if q.Len() == 0 {
					nout = nil
					continue
				}
				next = q.Dequeue()
				continue
			default:
			}
			select {
			case v, ok := <-in:
				if !ok {
					in = nil
					continue
				}
				if nout == nil {
					nout = out
					next = v
				} else {
					q.Enqueue(v)
				}
			case nout <- next:
				if q.Len() == 0 {
					nout = nil
					continue
				}
				next = q.Dequeue()
			}
		}
	}()
	return out
}

type SimpleQueue[T any] struct {
	data []T
	cap  int
	mu   sync.Mutex
}

func NewSimpleQueue[T any](capacity int) *SimpleQueue[T] {
	return &SimpleQueue[T]{data: make([]T, 0, capacity), cap: capacity}
}

func (q *SimpleQueue[T]) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.data)
}

func (q *SimpleQueue[T]) Cap() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.cap
}

func (q *SimpleQueue[T]) SetCap(v int) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if v < len(q.data) {
		q.data = q.data[:v]
	}
	q.cap = v
}

func (q *SimpleQueue[T]) Dequeue() (t T) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.data) == 0 {
		var zero T
		return zero
	}
	t = q.data[0]
	q.data = q.data[1:]
	return t
}

func (q *SimpleQueue[T]) Enqueue(t T) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.data) < q.cap {
		q.data = append(q.data, t)
	}
}

func TestBufLongLivedCustomQueue(t *testing.T) {
	in := make(chan int)
	q := NewSimpleQueue[int](10)
	out := BufLongLivedCustomQueue(in, q)

	go func() {
		for i := 0; i < 20; i++ {
			in <- i
		}
		close(in)
	}()

	expected := 0
	for v := range out {
		if v != expected {
			t.Errorf("expected %d, got %d", expected, v)
		}
		expected++
	}

	if expected != 20 {
		t.Errorf("expected to process 20 items, got %d", expected)
	}
}
