package bowl

func (b *Bowl[T]) SetClean(clean func(o T, a ...any)) {
	b.clean = clean
}
