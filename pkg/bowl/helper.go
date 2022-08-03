package bowl

import "sync/atomic"

func (b *Bowl[T]) incrementSize(i int32) {
	atomic.AddInt32(&b.size, i)
}

func (b *Bowl[T]) decrementSize(i int32) {
	atomic.AddInt32(&b.size, i)
}
