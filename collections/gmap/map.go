// Package gmap is a wrapper of golang map.
package gmap

type Map[KEY comparable, VALUE any] interface {
	Contains(key KEY) bool
	Put(key KEY, value VALUE)
	Delete(key KEY) VALUE
	PutIfAbsent(key KEY, value VALUE) bool
	DeleteIfPresent(key KEY) (VALUE, bool)
	Keys() []KEY
	Values() []VALUE
	Size() int

	ToBuiltIn() map[KEY]VALUE
}

type gmap[KEY comparable, VALUE any] struct {
	data map[KEY]VALUE
}

func New[KEY comparable, VALUE any]() Map[KEY, VALUE] {
	return &gmap[KEY, VALUE]{
		data: make(map[KEY]VALUE),
	}
}

func (m *gmap[KEY, VALUE]) Contains(key KEY) bool {
	_, ok := m.data[key]
	return ok
}

func (m *gmap[KEY, VALUE]) Put(key KEY, value VALUE) {
	m.data[key] = value
}

func (m *gmap[KEY, VALUE]) Delete(key KEY) VALUE {
	val := m.data[key]
	delete(m.data, key)
	return val
}

func (m *gmap[KEY, VALUE]) PutIfAbsent(key KEY, value VALUE) bool {
	if _, ok := m.data[key]; !ok {
		m.data[key] = value
		return true
	}
	return false
}

func (m *gmap[KEY, VALUE]) DeleteIfPresent(key KEY) (val VALUE, has bool) {
	has = false
	if _, ok := m.data[key]; ok {
		val = m.data[key]
		has = true
	}
	return
}

func (m *gmap[KEY, VALUE]) Keys() []KEY {
	result := make([]KEY, 0, m.Size())
	for key := range m.data {
		result = append(result, key)
	}
	return result
}

func (m *gmap[KEY, VALUE]) Values() []VALUE {
	result := make([]VALUE, m.Size())
	for key := range m.data {
		result = append(result, m.data[key])
	}
	return result
}

func (m *gmap[KEY, VALUE]) Size() int {
	return len(m.data)
}

func (m *gmap[KEY, VALUE]) ToBuiltIn() map[KEY]VALUE {
	result := make(map[KEY]VALUE)
	for k, v := range m.data {
		result[k] = v
	}
	return result
}
