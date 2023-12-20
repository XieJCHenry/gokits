package ordered_map

import (
	"github.com/XieJCHenry/gokits/collections/gmap"
	"github.com/XieJCHenry/gokits/collections/slice"
)

type OrderedMap[KEY comparable, VALUE any] interface {
	gmap.Map[KEY, VALUE]
	At(index int) VALUE
}

type valueWrapper[VALUE any] struct {
	val   VALUE
	index int
}

type orderedMap[KEY comparable, VALUE any] struct {
	keyList slice.Slice[KEY]
	data    gmap.Map[KEY, *valueWrapper[VALUE]]
}

func New[KEY comparable, VALUE any]() OrderedMap[KEY, VALUE] {
	return &orderedMap[KEY, VALUE]{
		data:    gmap.New[KEY, *valueWrapper[VALUE]](),
		keyList: slice.New[KEY](),
	}
}

func (o *orderedMap[KEY, VALUE]) Contains(key KEY) bool {
	return o.data.Contains(key)
}

func (o *orderedMap[KEY, VALUE]) Put(key KEY, value VALUE) {
	index := o.keyList.Size()
	o.keyList.Append(key)
	o.data.Put(key, &valueWrapper[VALUE]{
		val:   value,
		index: index,
	})
}

func (o *orderedMap[KEY, VALUE]) Get(key KEY) (val VALUE) {
	wrapper := o.data.Get(key)
	if wrapper != nil {
		val = wrapper.val
	}
	return val
}

func (o *orderedMap[KEY, VALUE]) GetOrDefault(key KEY, defaultValue VALUE) (val VALUE) {
	wrapper := o.data.GetOrDefault(key, nil)
	if wrapper != nil {
		val = wrapper.val
	} else {
		val = defaultValue
	}
	return val
}

// Delete is not O(1) operation
func (o *orderedMap[KEY, VALUE]) Delete(key KEY) (val VALUE) {
	if wrapper, ok := o.data.DeleteIfPresent(key); ok {
		index := wrapper.index
		o.keyList.RemoveAt(index)
		val = wrapper.val
	}
	return val
}

func (o *orderedMap[KEY, VALUE]) PutIfAbsent(key KEY, value VALUE) bool {
	var absent bool
	if !o.data.Contains(key) {
		absent = true
		index := o.keyList.Size()
		o.keyList.Append(key)
		o.data.Put(key, &valueWrapper[VALUE]{
			val:   value,
			index: index,
		})
	}
	return absent
}

func (o *orderedMap[KEY, VALUE]) DeleteIfPresent(key KEY) (val VALUE, present bool) {
	var wrapper *valueWrapper[VALUE]
	if wrapper, present = o.data.DeleteIfPresent(key); present {
		index := wrapper.index
		o.keyList.RemoveAt(index)
		val = wrapper.val
	}
	return
}

// Keys is ordered
func (o *orderedMap[KEY, VALUE]) Keys() []KEY {
	return o.keyList.ToBuiltIn()
}

// Values is ordered by keys
func (o *orderedMap[KEY, VALUE]) Values() []VALUE {
	result := make([]VALUE, 0, o.keyList.Size())
	o.keyList.ForEach(func(index int, key KEY, result any) {
		res := result.([]VALUE)
		val := o.data.Get(key).val
		res = append(res, val)
	}, result)
	return result
}

func (o *orderedMap[KEY, VALUE]) Size() int {
	return o.data.Size()
}

func (o *orderedMap[KEY, VALUE]) ToBuiltIn() map[KEY]VALUE {
	result := make(map[KEY]VALUE)
	o.data.ForEach(func(key KEY, val *valueWrapper[VALUE], result any) {
		res := result.(map[KEY]VALUE)
		res[key] = val.val
	}, result)
	return result
}

func (o *orderedMap[KEY, VALUE]) At(index int) (val VALUE) {
	if index >= 0 && index < o.keyList.Size() {
		val = o.data.Get(o.keyList.At(index)).val
	}
	return val
}

func (o *orderedMap[KEY, VALUE]) ForEach(each func(key KEY, val VALUE, result any), result any) {
	keysForEach := func(index int, x KEY, result0 any) {
		val := o.data.Get(x).val
		each(x, val, result)
	}
	o.keyList.ForEach(keysForEach, nil)
}
