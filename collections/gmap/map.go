// Package gmap is a wrapper of golang map.
package gmap

type Map[KEY comparable, VALUE any] interface {
	Contains(key KEY) bool
	Put(key KEY, value VALUE)
	Get(key KEY) VALUE
	GetOrDefault(key KEY, defaultValue VALUE) VALUE
	Delete(key KEY) VALUE
	PutIfAbsent(key KEY, value VALUE) bool
	DeleteIfPresent(key KEY) (VALUE, bool)
	Keys() []KEY
	Values() []VALUE
	Size() int

	ToBuiltIn() map[KEY]VALUE
	ForEach(each func(key KEY, val VALUE, result any), result any)
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

func (m *gmap[KEY, VALUE]) Get(key KEY) VALUE {
	return m.data[key]
}

func (m *gmap[KEY, VALUE]) GetOrDefault(key KEY, defaultValue VALUE) VALUE {
	var val VALUE
	var ok bool
	if val, ok = m.data[key]; !ok {
		val = defaultValue
	}
	return val
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
	if val, has = m.data[key]; has {
		delete(m.data, key)
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

func (m *gmap[KEY, VALUE]) ForEach(each func(key KEY, val VALUE, result any), result any) {
	for k, v := range m.data {
		each(k, v, result)
	}
}
