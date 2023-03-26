// Package slices contains some helper functions for built-in slice
package slices

import (
	"github.com/XieJCHenry/gokits/collections/set"
	"github.com/XieJCHenry/gokits/collections/tuple"
	"math"
)

func Contains(arr []interface{}, x interface{}) bool {
	return IndexOf(arr, x) != -1
}

func IndexOf(arr []interface{}, x interface{}) int {
	idx := -1
	for i := range arr {
		if arr[i] == x {
			idx = i
			break
		}
	}
	return idx
}

func Remove(arr []interface{}, x interface{}) []interface{} {
	idx := IndexOf(arr, x)
	if idx != -1 {
		arr = append(arr[0:idx], arr[idx+1:])
	}
	return arr
}

func RemoveAt(arr []interface{}, index int) ([]interface{}, interface{}) {
	if index < 0 || index >= len(arr) {
		return arr, nil
	}
	x := arr[index]
	arr = Remove(arr, x)
	return arr, x
}

func ToSet[T comparable](arr []T) set.Set[T] {
	return set.NewFrom[T](arr...)
}

func ToTupleList(arr1 []interface{}, arr2 []interface{}) []tuple.Tuple {
	minSize := int(math.Min(float64(len(arr1)), float64(len(arr2))))
	result := make([]tuple.Tuple, 0, minSize)
	for i := 0; i < minSize; i++ {
		t := tuple.New(2)
		t.Set(0, arr1[i])
		t.Set(1, arr2[i])
		result[i] = t
	}
	return result
}

func ToMap[T comparable](keys []T, vals []interface{}) map[T]interface{} {
	result := make(map[T]interface{})
	minSize := int(math.Min(float64(len(keys)), float64(len(vals))))

	for i := 0; i < minSize; i++ {
		key := keys[i]
		val := vals[i]
		// 仅不存在时插入，存在重复时略过
		if _, ok := result[key]; !ok {
			result[key] = val
		}
	}
	return result
}
