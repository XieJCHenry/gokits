package slice

type Slice[T comparable] interface {
	At(index int) T
	Append(x T)
	AppendIfAbsent(x T) bool
	Contains(x T) bool
	IndexOf(x T) int
	InsertAt(index int, x T)
	Remove(x T)
	RemoveAt(index int) T
	RemoveIfPresent(x T) bool
	Size() int

	ToBuiltIn() []T
	Filter(filter func(x T) bool) Slice[T]
}

type slice[T comparable] struct {
	data []T
}

func New[T comparable]() Slice[T] {
	return &slice[T]{
		data: make([]T, 0),
	}
}

func (s *slice[T]) At(index int) (val T) {
	if index >= 0 && index < len(s.data) {
		val = s.data[index]
	}
	return
}

func (s *slice[T]) Append(x T) {
	s.data = append(s.data, x)
}

func (s *slice[T]) AppendIfAbsent(x T) bool {
	if !s.Contains(x) {
		s.Append(x)
		return true
	}
	return false
}

func (s *slice[T]) Contains(x T) bool {
	return s.IndexOf(x) != -1
}

func (s *slice[T]) IndexOf(x T) int {
	for index := range s.data {
		if s.data[index] == x {
			return index
		}
	}
	return -1
}

func (s *slice[T]) InsertAt(index int, x T) {
	if index >= len(s.data)-1 {
		s.data = append(s.data, x)
	} else if index >= 0 && index < len(s.data)-1 {
		preSlice := append(s.data[:index], x)
		s.data = append(preSlice, s.data[index:]...)
	}
}

func (s *slice[T]) Remove(x T) {
	index := s.IndexOf(x)
	s.RemoveAt(index)
}

func (s *slice[T]) RemoveAt(index int) (val T) {
	if index >= 0 && index < s.Size() {
		val = s.At(index)

		if index == len(s.data)-1 {
			s.data = s.data[:len(s.data)-1]
		} else if index != -1 {
			s.data = append(s.data[:index], s.data[:index+1]...)
		}
	}
	return
}

func (s *slice[T]) RemoveIfPresent(x T) bool {
	index := s.IndexOf(x)
	if index != -1 {
		s.RemoveAt(index)
	}
	return index != -1
}

func (s *slice[T]) Size() int {
	return len(s.data)
}

func (s *slice[T]) Filter(filter func(x T) bool) Slice[T] {
	newSlice := New[T]()
	for i := range s.data {
		if filter(s.data[i]) {
			newSlice.Append(s.data[i])
		}
	}
	return newSlice
}

func (s *slice[T]) ToBuiltIn() []T {
	result := make([]T, s.Size())
	copy(result, s.data)
	return result
}
