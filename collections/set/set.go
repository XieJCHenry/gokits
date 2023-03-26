package set

type Set[T comparable] interface {
	Contains(x T) bool
	Size() int

	InsertIfAbsent(x T) bool
	RemoveIfPresent(x T) bool
	ToArray() []T
	ToBuiltIn() map[T]struct{}
}

type set[T comparable] struct {
	data map[T]struct{}
}

func New[T comparable]() Set[T] {
	return &set[T]{
		data: make(map[T]struct{}),
	}
}

func NewFrom[T comparable](elems ...T) Set[T] {
	s := &set[T]{
		data: make(map[T]struct{}),
	}
	for i := range elems {
		s.data[elems[i]] = struct{}{}
	}
	return s
}

func (s *set[T]) Contains(x T) bool {
	_, ok := s.data[x]
	return ok
}

func (s *set[T]) Size() int {
	return len(s.data)
}

func (s *set[T]) InsertIfAbsent(x T) bool {
	if _, ok := s.data[x]; !ok {
		s.data[x] = struct{}{}
		return true
	}
	return false
}

func (s *set[T]) RemoveIfPresent(x T) bool {
	if _, ok := s.data[x]; ok {
		delete(s.data, x)
		return true
	}
	return false
}

func (s *set[T]) ToArray() []T {
	result := make([]T, 0, s.Size())
	for x := range s.data {
		result = append(result, x)
	}
	return result
}

func (s *set[T]) ToBuiltIn() map[T]struct{} {
	result := make(map[T]struct{})
	for x := range s.data {
		result[x] = struct{}{}
	}
	return s.data
}
