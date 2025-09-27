package day24

// QUEUE
type Queue[T any] []T

func NewQueue[T any]() Queue[T] {
	return Queue[T]{}
}

func NewQueueWith[T any](item T) Queue[T] {
	return Queue[T]{item}
}

func (q *Queue[T]) Pop() T {
	removed := (*q)[0]
	(*q) = (*q)[1:]
	return removed
}
func (q *Queue[T]) Push(c T) {
	*q = append(*q, c)
}
