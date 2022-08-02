package bowl

import (
	"sync/atomic"
)

var DefaultMax int32 = 5

type Bowler[T any] interface {
	Get() T
	Return(T)
}

type Bowl[T any] struct {
	// pool is containg all the objects
	pool chan T

	// new is the function to use
	// to retrieve a new object
	new func() T

	// size is the maximum number
	// of elements that the pool can
	// contain. The default is 5
	max int32

	// size is a approximiation
	// of the elements that are
	// in the pool at the point of
	// function call
	size int32
}

func New[T any](max int32, new func() T) Bowl[T] {

	if max == 0 {
		max = DefaultMax
	}

	b := Bowl[T]{
		pool: make(chan T, max),
		max:  max,
		new:  new,
	}

	return b

}

func (b *Bowl[T]) Return(o T) {

	if len(b.pool) == int(b.max) {
		return
	}

	b.incrementSize(1)
	b.pool <- o

}

func (b *Bowl[T]) Get() T {
	if len(b.pool) > 0 {
		b.decrementSize(-1)
		return <-b.pool
	}

	return b.new()
}

func (b *Bowl[T]) Size() int32 {
	return b.size
}

func (b *Bowl[T]) Max() int32 {
	return b.max
}

func (b *Bowl[T]) incrementSize(i int32) {
	atomic.AddInt32(&b.size, i)
}

func (b *Bowl[T]) decrementSize(i int32) {
	atomic.AddInt32(&b.size, i)
}
