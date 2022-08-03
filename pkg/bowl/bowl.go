package bowl

var DefaultLimit int32 = 5

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

	// clean will clean the object.
	// Default will be an empty function.
	clean func(*T, ...any)

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

// New Will return a new Bowl. The default limit
// is 5 but can be changed by reassigning `DefaultLimit`.
// `new` is the factory function that will be used
// to create a new Element `T` if there are no elements
// in the bowl.
func New[T any](max int32, new func() T) Bowl[T] {

	if max == 0 {
		max = DefaultLimit
	}

	b := Bowl[T]{
		pool:  make(chan T, max),
		max:   max,
		new:   new,
		clean: func(o *T, args ...any) {},
	}

	return b

}

// Return will put the object `o` back into
// the bowl. It will always clean `o` using
// the provided clean function  and passing the
// args to the clean function.
func (b *Bowl[T]) Return(o T, args ...any) {
	if len(b.pool) == int(b.max) {
		return
	}

	b.clean(&o, args)
	b.incrementSize(1)
	b.pool <- o

}

// Get is retrieving `T` from the bowl. If
// the bowl is empty, it will create a new
// element using the `new` function provided
// while creating the bowl.
func (b *Bowl[T]) Get() T {
	if len(b.pool) > 0 {
		b.decrementSize(-1)
		return <-b.pool
	}

	return b.new()
}
