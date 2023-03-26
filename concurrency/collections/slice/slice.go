package slice

import (
	"sync"

	"github.com/XieJCHenry/gokits/collections/slice"
)

type sync_slice[T comparable] struct {
	mtx  sync.Mutex
	data slice.Slice[T]
}

func New[T comparable]() slice.Slice[T] {
	return &sync_slice[T]{
		mtx:  sync.Mutex{},
		data: slice.New[T](),
	}
}

func (ss *sync_slice[T]) At(index int) T {
	ss.mtx.Lock()
	defer ss.mtx.Unlock()

	return ss.data.At(index)
}

func (ss *sync_slice[T]) Append(x T) {
	ss.mtx.Lock()
	defer ss.mtx.Unlock()
}

func (ss *sync_slice[T]) AppendIfAbsent(x T) bool {
	ss.mtx.Lock()
	defer ss.mtx.Unlock()

	return ss.data.AppendIfAbsent(x)
}

func (ss *sync_slice[T]) Contains(x T) bool {
	ss.mtx.Lock()
	defer ss.mtx.Unlock()

	return ss.data.Contains(x)
}

func (ss *sync_slice[T]) IndexOf(x T) int {
	ss.mtx.Lock()
	defer ss.mtx.Unlock()

	return ss.data.IndexOf(x)
}

func (ss *sync_slice[T]) InsertAt(index int, x T) {
	ss.mtx.Lock()
	defer ss.mtx.Unlock()
}

func (ss *sync_slice[T]) Remove(x T) {
	ss.mtx.Lock()
	defer ss.mtx.Unlock()
}

func (ss *sync_slice[T]) RemoveAt(index int) T {
	ss.mtx.Lock()
	defer ss.mtx.Unlock()

	return ss.data.RemoveAt(index)
}

func (ss *sync_slice[T]) RemoveIfPresent(x T) bool {
	ss.mtx.Lock()
	defer ss.mtx.Unlock()

	return ss.data.RemoveIfPresent(x)
}

func (ss *sync_slice[T]) Size() int {
	ss.mtx.Lock()
	defer ss.mtx.Unlock()

	return ss.data.Size()
}

func (ss *sync_slice[T]) ToBuiltIn() []T {
	ss.mtx.Lock()
	defer ss.mtx.Unlock()

	return ss.data.ToBuiltIn()
}

func (ss *sync_slice[T]) Filter(filter func(x T) bool) slice.Slice[T] {
	ss.mtx.Lock()
	defer ss.mtx.Unlock()

	return ss.data.Filter(filter)
}
