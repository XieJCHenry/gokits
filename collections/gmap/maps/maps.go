// Package maps support some helper functions for builtin map
package maps

import (
	"errors"
	"github.com/XieJCHenry/gokits/collections/tuple"
	"math"
)

func GetKeys[KEY comparable, VALUE any](m map[KEY]VALUE) []KEY {
	result := make([]KEY, 0, len(m))
	for key := range m {
		result = append(result, key)
	}
	return result
}

func GetValues[KEY comparable, VALUE any](m map[KEY]VALUE) []VALUE {
	result := make([]VALUE, 0, len(m))
	for _, val := range m {
		result = append(result, val)
	}
	return result
}

func NewFromTuples[KEY comparable, VALUE any](tuples []tuple.Tuple) (map[KEY]VALUE, error) {
	result := make(map[KEY]VALUE)
	if len(tuples) == 0 {
		return result, nil
	}
	if tuples[0].Size() != 2 {
		return nil, errors.New("size of tuple required 2")
	}

	for _, t := range tuples {
		key0 := t.At(0)
		value := t.At(1)
		if key, ok := key0.(KEY); ok {
			result[key] = value.(VALUE)
		}
	}
	return result, nil
}

func NewFromStruct[KEY comparable, VALUE any](elements []VALUE, keyGenerator func(elem VALUE) (KEY, bool)) map[KEY]VALUE {
	result := make(map[KEY]VALUE)
	for _, elem := range elements {
		if key, ok := keyGenerator(elem); ok {
			result[key] = elem
		}
	}
	return result
}

func NewFromKeyValues[KEY comparable, VALUE any](keys []KEY, values []VALUE) map[KEY]VALUE {
	result := make(map[KEY]VALUE)
	if len(keys) == 0 {
		return result
	}
	minSize := int(math.Min(float64(len(keys)), float64(len(values))))
	for i := 0; i < minSize; i++ {
		result[keys[i]] = values[i]
	}
	return result
}

func NewFromSlice[T any, U comparable, V any](data []T, convertFunc func(d T) (U, V)) map[U]V {
	result := make(map[U]V)
	if len(data) > 0 {
		for i := range data {
			d := data[i]
			key, val := convertFunc(d)
			result[key] = val
		}
	}

	return result
}
