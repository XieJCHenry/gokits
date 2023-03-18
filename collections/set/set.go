package set

import "sort"

type Set interface {
	Contains(x interface{}) bool
	Size() int

	InsertIfAbsent(x interface{}) bool
	RemoveIfPresent(x interface{}) bool
	ToArray() []interface{}
	ToOrderedArray(cmp func(i, j int) bool) []interface{}
}

type set[T comparable] struct {
	data map[T]struct{}
}

func New[T comparable]() Set {
	return &set[T]{
		data: make(map[T]struct{}),
	}
}

func NewFrom[T comparable](elems ...T) Set {
	s := &set[T]{
		data: make(map[T]struct{}),
	}
	for i := range elems {
		s.data[elems[i]] = struct{}{}
	}
	return s
}

func (s *set[T]) Contains(x interface{}) bool {
	_, ok := s.data[x]
	return ok
}

func (s *set[T]) Size() int {
	return len(s.data)
}

func (s *set[T]) InsertIfAbsent(x interface{}) bool {
	if _, ok := s.data[x]; !ok {
		s.data[x] = struct{}{}
		return true
	}
	return false
}

func (s *set[T]) RemoveIfPresent(x interface{}) bool {
	if _, ok := s.data[x]; ok {
		delete(s.data, x)
		return true
	}
	return false
}

func (s *set[T]) ToArray() []interface{} {
	result := make([]interface{}, 0, s.Size())
	for x := range s.data {
		result = append(result, x)
	}
	return result
}

func (s *set[T]) ToOrderedArray(cmp func(i, j int) bool) []interface{} {
	result := s.ToArray()
	sort.Slice(result, cmp)
	return result
}
