package ordered_map

import (
	"github.com/XieJCHenry/gokits/collections/gmap"
	"github.com/XieJCHenry/gokits/collections/slice"
)

type OrderedMap[KEY comparable, VALUE any] interface {
	gmap.Map[KEY, VALUE]

	OrderedKeys() slice.Slice[KEY]
}

type orderedMap[KEY comparable, VALUE any] struct {
	keys    slice.Slice[KEY]
	dataMap gmap.Map[KEY, VALUE]
}

func New[KEY comparable, VALUE any]() OrderedMap[KEY, VALUE] {
	return &orderedMap[KEY, VALUE]{
		keys:    slice.New[KEY](),
		dataMap: gmap.New[KEY, VALUE](),
	}
}

func (o *orderedMap[KEY, VALUE]) Contains(key KEY) bool {
	return o.dataMap.Contains(key)
}

func (o *orderedMap[KEY, VALUE]) Put(key KEY, value VALUE) {
	if o.dataMap.PutIfAbsent(key, value) {
		o.keys.Append(key)
	}
}

func (o *orderedMap[KEY, VALUE]) Get(key KEY) VALUE {
	return o.dataMap.Get(key)
}

func (o *orderedMap[KEY, VALUE]) GetOrDefault(key KEY, defaultValue VALUE) VALUE {
	return o.dataMap.GetOrDefault(key, defaultValue)
}

func (o *orderedMap[KEY, VALUE]) Delete(key KEY) VALUE {
	existVal := o.dataMap.Delete(key)
	o.keys.Remove(key)
	return existVal
}

func (o *orderedMap[KEY, VALUE]) PutIfAbsent(key KEY, value VALUE) bool {
	put := o.dataMap.PutIfAbsent(key, value)
	if put {
		o.keys.Append(key)
	}
	return put
}

func (o *orderedMap[KEY, VALUE]) DeleteIfPresent(key KEY) (VALUE, bool) {
	existVal, deleted := o.dataMap.DeleteIfPresent(key)
	if deleted {
		o.keys.Remove(key)
	}
	return existVal, deleted
}

func (o *orderedMap[KEY, VALUE]) Keys() []KEY {
	return o.dataMap.Keys()
}

func (o *orderedMap[KEY, VALUE]) Values() []VALUE {
	return o.dataMap.Values()
}

func (o *orderedMap[KEY, VALUE]) Size() int {
	return o.dataMap.Size()
}

func (o *orderedMap[KEY, VALUE]) ToBuiltIn() map[KEY]VALUE {
	return o.dataMap.ToBuiltIn()
}

func (o *orderedMap[KEY, VALUE]) OrderedKeys() slice.Slice[KEY] {
	return o.keys
}
