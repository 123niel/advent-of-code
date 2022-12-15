package misc

type Stack[T any] struct {
	Items []T
}

func CopyStacks[T any](stacks []Stack[T]) []Stack[T] {
	copy := make([]Stack[T], len(stacks))

	for i := 0; i < len(stacks); i++ {
		copy[i] = *stacks[i].copy()
	}

	return copy
}

func (s *Stack[T]) copy() *Stack[T] {
	var copy Stack[T]
	copy.Items = s.Items
	return &copy
}

func (s *Stack[T]) zero() T {
	var zero T
	return zero
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.Items) == 0
}

func (s *Stack[T]) Push(val T) {
	s.Items = append(s.Items, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		return s.zero(), false
	} else {
		i := len(s.Items) - 1
		element := s.Items[i]
		s.Items = s.Items[:i]
		return element, true
	}
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		return s.zero(), false
	} else {
		return s.Items[len(s.Items)-1], true
	}
}

func (s *Stack[T]) PopMutliple(n int) ([]T, bool) {
	var len = len(s.Items)

	if len < n {
		return []T{}, false
	} else {
		elements := s.Items[len-n : len]
		s.Items = s.Items[:len-n]
		return elements, true
	}
}

func (s *Stack[T]) PushMultiple(vals []T) {
	s.Items = append(s.Items, vals...)
}
