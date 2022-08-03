package bowl

// Size returns the quantity of elements
// that are in the pool.
func (b *Bowl[T]) Size() int32 {
	return b.size
}

// Max is the maximum quantity
// the pool can hold.
func (b *Bowl[T]) Max() int32 {
	return b.max
}
