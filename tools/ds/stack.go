package ds

import "errors"

type Stack[T any] []T

// Errors
var ErrEmptyStack = errors.New("stack is empty")

func (s *Stack[T]) Move(amount int, to *Stack[T], preserveOrder bool) {
	var items []T

	for range amount {
		item, err := (*s).Pop()
		if err == nil {
			items = append(items, item)

		}
	}
	if preserveOrder {
		for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
			items[i], items[j] = items[j], items[i]
		}
	}
	for _, item := range items {
		to.Push(item)
	}
}

func (s *Stack[T]) Pop() (T, error) {
	if len(*s) == 0 {
		var blank T
		return blank, ErrEmptyStack
	}
	removed := (*s)[0]
	*s = (*s)[1:]
	return removed, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if len(*s) == 0 {
		var blank T
		return blank, ErrEmptyStack
	}
	return (*s)[0], nil
}

func (s *Stack[T]) Push(item T) {
	*s = append([]T{item}, *s...)
}
